# About Project

This project is a backend application for managing a hotel.

## Usage

To run the application:

1. Clone the repository
2. Install dependencies: `go mod tidy`
3. Set up your PostgreSQL database and update the configuration in `config/.env`
4. Run the application: `go run main.go`
   
The server will start on `localhost:8080` by default.

## Available Endpoints

### Users
- `POST /user/login` - Login user
- `POST /user/register` - Create a new user

### Rooms
- `GET /api/rooms` - Get a list of all rooms
- `POST /api/rooms` - Add a new room
- `GET /api/rooms/:id` - Get room details
- `PUT /api/rooms/:id` - Update room information
- `DELETE /api/rooms/:id` - Remove a room

### Bookings
- `GET /api/bookings` - Get a list of all Bookings
- `POST /api/bookings` - Create a new booking 
- `GET /api/bookings/:id` - Get booking details
- `PUT /api/bookings/:id` - Update booking information
- `DELETE /api/bookings/:id` - Remove a booking

### Categories
- `GET /api/payments` - Get a list of all payments
- `POST /api/payments` - Create a new payments
- `GET /api/payments/:id` - Get payment details
- `PUT /api/payments/:id` - Update payment information
- `DELETE /api/payments/:id` - Delete a payment
  
## Configuration

The application configuration is stored in `config/.env`. Make sure to update the database connection details before running the application.

## Template For Adding Data
### For Adding data user / register
{
    "username":"salman",
    "password":"salman123",
    "email":"salman@gmail.com",
    "role":"user"
}

### For adding room data
{
    "id": 1,
    "room_number": "101",
    "type": "Single",
    "price": 100,
    "status": "Available"
}
### For adding booking data
{
    "id": 1,
    "room_id": 1,
    "user_id": 1,
    "check_in": "2024-08-10T12:00:00Z",
    "check_out": "2024-08-12T12:00:00Z",
    "status": "Confirmed"
}
### For adding payment data
{
  "id": 1,
  "booking_id": 1,
  "amount": 200,
  "payment_date": "2024-08-10T12:00:00Z",
  "status": "Completed"
}
