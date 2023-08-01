# Music Streaming API

This is a music streaming API that allows you to manage songs.

## Getting Started

Follow the steps below to set up and run the web service.

### Prerequisites

- Go programming language (version 1.15 or higher)
- Git

### Installation

1. Clone the repository:

   ```
   git clone https://github.com/mucha-fauzy/musics-streaming-api.git
   ```

2. Install dependencies:

   ```
   go mod download
   ```

### Generate Swagger Documentation

To generate the Swagger documentation, run the following command:

```
swag init
```

### How to Run

1. Start the web service:

   ```
   go run main.go
   ```

   The server will start running at `http://localhost:8080`.

2. Access Swagger UI:

   Open your web browser and visit `http://localhost:8080/swagger/index.html`. This will take you to the Swagger UI, where you can view and test the API endpoints.

### API Endpoints

- **POST** `/api/v1/songs`: Create a new song.
- **GET** `/api/v1/songs`: Get a list of all songs. You can use query parameters to filter, sort, and paginate the results.
- **PUT** `/api/v1/songs/{id}`: Update an existing song by its ID.
- **DELETE** `/api/v1/songs/{id}`: Delete a song by its ID.

### Sample Data
The API comes with no sample data. You can use POST to generate some data and use others provided endpoints to interact with the data.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to create a pull request or submit an issue.

## Authors

- Muchamad Fauzy (muchamad.fauzy@evermos.com)

## Acknowledgments

- [Swaggo](https://github.com/swaggo/swag) - For generating Swagger documentation.