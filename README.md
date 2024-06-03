# Exoplanet Microservice

This microservice allows users to manage a catalog of exoplanets, with functionalities to add, list, retrieve, update, and delete exoplanets. Additionally, it can estimate the fuel required for a trip to a specific exoplanet.

## Endpoints

- `POST /exoplanets` - Add a new exoplanet
- `GET /exoplanets` - List all exoplanets
- `GET /exoplanets/{id}` - Get exoplanet details by ID
- `PUT /exoplanets/{id}` - Update an existing exoplanet
- `DELETE /exoplanets/{id}` - Delete an exoplanet
- `GET /exoplanets/{id}/fuel` - Get fuel estimation for a trip to an exoplanet

## Running Locally

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/exoplanet-service.git
   cd exoplanet-service
