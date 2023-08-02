package handlers

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"musics_streaming_api/internal/models"
	"musics_streaming_api/internal/services"

	"github.com/go-chi/chi"
)

type Handler struct {
	service services.Service
}

func ProvideHandler(service services.Service) *Handler {
	return &Handler{service: service}
}

// @Summary Create a new song
// @Description Create a new song with the given details
// @Produce json
// @Param song body models.Songs true "New song details"
// @Success 201 {object} models.Songs
// @Failure 400 {object} ErrorResponse
// @Router /api/v1/songs [post]
func (h *Handler) CreateSongHandler(w http.ResponseWriter, r *http.Request) {
	var newSong models.Songs
	err := json.NewDecoder(r.Body).Decode(&newSong)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	song, success := h.service.AddSong(&newSong)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create a song entry"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
		"message": "Song entry successfully created",
		"song":    song,
	}
	json.NewEncoder(w).Encode(response)
}

// @Summary Get all songs
// @Description Get a list of all songs with optional filtering, sorting, and pagination
// @Produce json
// @Param title query string false "Filter by song title"
// @Param sort_by query string false "Sort by 'release_date' or 'duration'"
// @Param order query string false "Order by 'asc' or 'desc'"
// @Param page query int false "Page number (default 1)"
// @Param size query int false "Number of songs per page (default 5)"
// @Success 200 {array} models.Songs
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/songs [get]
func (h *Handler) GetSongHandler(w http.ResponseWriter, r *http.Request) {
	// Get query parameters
	filterTitle := r.URL.Query().Get("title")
	sortBy := r.URL.Query().Get("sort_by")
	order := r.URL.Query().Get("order")
	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	// Convert page and size parameters to integers
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil || size < 1 {
		size = 5
	}

	// Get all songs from the service
	songs := h.service.GetAllSongs()

	// Filter songs by title if provided
	if filterTitle != "" {
		filteredSongs := make([]*models.Songs, 0)
		for _, song := range songs {
			if strings.Contains(strings.ToLower(song.Title), strings.ToLower(filterTitle)) {
				filteredSongs = append(filteredSongs, song)
			}
		}
		songs = filteredSongs
	}

	// Sort songs by release date or duration if provided
	if sortBy == "release_date" {
		sort.Slice(songs, func(i, j int) bool {
			dateI, err := time.Parse("2006-01-02", songs[i].ReleaseDate)
			if err != nil {
				return false
			}
			dateJ, err := time.Parse("2006-01-02", songs[j].ReleaseDate)
			if err != nil {
				return false
			}
			return dateI.Before(dateJ)
		})
	} else if sortBy == "duration" {
		sort.Slice(songs, func(i, j int) bool {
			return songs[i].Duration < songs[j].Duration
		})
	}

	// Order songs in ascending or descending order if provided
	if order == "desc" {
		reverseSlice(songs)
	}

	// Paginate the songs based on the page and size
	startIdx := (page - 1) * size
	endIdx := startIdx + size
	if startIdx >= len(songs) {
		songs = []*models.Songs{}
	} else if endIdx > len(songs) {
		songs = songs[startIdx:]
	} else {
		songs = songs[startIdx:endIdx]
	}

	// Check if there are any songs after filtering and sorting
	if len(songs) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Songs not found"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(songs)
}

func reverseSlice(s []*models.Songs) {
	for i := len(s)/2 - 1; i >= 0; i-- {
		opp := len(s) - 1 - i
		s[i], s[opp] = s[opp], s[i]
	}
}

// @Summary Update an existing song
// @Description Update an existing song with the given details
// @Produce json
// @Param id path string true "Song ID"
// @Param song body models.UpdateSongs true "Updated song details"
// @Success 200 {object} models.Songs
// @Failure 400 {object} ErrorResponse
// @Router /api/v1/songs/{id} [put]
func (h *Handler) UpdateSongHandler(w http.ResponseWriter, r *http.Request) {
	songID := chi.URLParam(r, "id")

	var updatedSong models.UpdateSongs
	err := json.NewDecoder(r.Body).Decode(&updatedSong)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	song, success := h.service.UpdateSong(songID, &updatedSong)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Song not found"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message": "Song successfully updated",
		"song":    song,
	}
	json.NewEncoder(w).Encode(response)
}

// @Summary Delete a song
// @Description Delete a song with the given ID
// @Produce json
// @Param id path string true "Song ID"
// @Success 204 "No Content"
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/songs/{id} [delete]
func (h *Handler) DeleteSongHandler(w http.ResponseWriter, r *http.Request) {
	songID := chi.URLParam(r, "id")

	success := h.service.DeleteSong(songID)
	if !success {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Song not found"})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// ErrorResponse is the response structure for error messages
type ErrorResponse struct {
	Error string `json:"error"`
}
