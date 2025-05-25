# Ticket Application (Mobile)



A full-stack ticket booking application built with Golang (backend) and React Native (mobile frontend). Built using **React Native**, **Golang**, **PostgreSQL**, **Docker**, and **TypeScript**. 

---

## Features

### Backend

- **User Authentication**: Secure registration and login with JWT tokens.
- **Event Management**: Create, update, delete, and list events.
- **Ticket Booking**: Book tickets for events, manage bookings, and view booking history.
- **Role-based Access**: Admin and user roles with appropriate permissions.
- **RESTful API**: Well-structured endpoints for all resources (fiber)
- **Database Integration**: Persistent storage using PostgreSQL.
- **Validation & Error Handling**: Robust input validation and descriptive error messages.
- **Scalable Architecture**: Modular codebase for easy extension.

### Mobile Frontend

- **User Registration & Login**: Simple onboarding and authentication.
- **Event Browsing**: List and search for available events.
- **Event Details**: View event information including date, time, location, and description.
- **Ticket Booking**: Book tickets directly from the app, including scanning functionality. 
- **Booking History**: View and manage your past and future bookings.
- **Responsive UI**: Clean, modern, and intuitive interface.
- **State Management**: Efficient state handling for smooth user experience.
- **Deployment**: Fully containerized using docker with a proper CI/CD pipeline and instructions. 

---

## Architecture

- **Backend**: Golang (Fiber), PostgreSQL, JWT authentication, Docker, CI/CD Pipeline
- **Frontend**: React Native, Redux (or Context API) for state management, REST API integration
- **Communication**: RESTful HTTP/JSON between mobile app and backend server

---

## Backend (Golang)

### Key Technologies

- Golang  
- Fiber
- JWT (authentication)  
- PostgreSQL (database)  

### Main Modules

- **User Module**: Handles registration, login, JWT token issuance, and user profile
- **Event Module**: CRUD operations for events
- **Booking Module**: Handles ticket booking, listing, and cancellation
- **Middleware**: Authentication and role-based access control

## Mobile Frontend (React Native)

### Key Technologies

- React Native  
- Redux or Context API  
- Axios (API calls)  
- React Navigation (navigation)  
- AsyncStorage (local storage)  

### Main Screens

- **Login/Register Screen**: User authentication  
- **Event List Screen**: Browse all available events  
- **Event Details Screen**: View detailed event info and book tickets  
- **Booking Screen**: View and manage your bookings  
- **Profile Screen**: View and edit user profile  

--- 