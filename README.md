# Subscription Management System

This repository provides a **scalable subscription management system** built with **Go** and **Firestore**. The system is designed to handle **secure authentication**, **real-time subscription monitoring**, **automated renewals**, and **efficient database handling**. 

## Features

- **User Authentication**: Secure login with JWT and Firebase Authentication.
- **Subscription Management**: Handle subscriptions, validations, renewals, and cancellations.
- **Real-Time Monitoring**: Monitor subscription changes in real-time with Firestore listeners.
- **Email Notifications**: Automatic notifications for subscription renewals and expirations.
- **Caching**: Temporary caching using Redis for faster data retrieval.
- **Database**: Scalable and efficient data storage with Firestore (or Postgres).
- **Testable**: Unit tests for critical functionalities.

## Technologies Used

- **Go**: Backend programming language.
- **Firestore**: NoSQL real-time database.
- **JWT**: For user authentication.
- **Redis**: For caching frequently accessed data.
- **SendGrid/SMTP**: For email notifications.
- **Postgres (optional)**: Relational database support.

## Project Structure

```plaintext
backend/
│── api/                   # ✅ Handles requests & business logic
│   ├── auth/              # 🔒 User authentication (JWT, Firebase Auth)
│   ├── subscription/      # 🔄 Subscription management
│   │   ├── validation.go  # ✅ Validates subscription data
│   │   ├── renewal.go     # ✅ Handles subscription renewals
│   │   ├── revoke.go      # ✅ Revokes expired subscriptions
│   │   ├── listener.go    # ✅ Firestore real-time monitoring
│── data/                  # ✅ Database layer (Firestore/Postgres)
│   ├── repository.go      # 🗂 Defines interface for subscription storage
│   ├── firestore.go       # 🔥 Firestore-specific implementation
│── config/                # ⚙️ Manages global configurations
│   ├── firebase.go        # 🔥 Firestore initialization
│   ├── settings.go        # ⚙️ App-wide settings
│── internal/              # 📦 Internal utilities (cache, notifications)
│   ├── cache/             # 🏎️ Temporary data caching (Redis)
│   ├── email/             # 📧 Handles notification emails
│── tests/                 # 🧪 Unit tests for major functions
│── main.go                # 🚀 API entry point

```
## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

- **Go**: Version 1.16+
- **Firestore**: Firebase project and Firestore database setup
- **Redis (Optional)**: For caching
- **SendGrid/SMTP**: Email service setup for notifications

### Installation

**1. Clone the repository:**

```bash
git clone https://github.com/yourusername/subscription-management-system.git
cd subscription-management-system
```
**2. Install dependencies:**

```bash
go mod tidy
```

**3. Set up Firebase Firestore:**

- *Follow the [Firebase Firestore setup guide.](https://firebase.google.com/docs/firestore)*

- *Create a Firebase project and enable Firestore.*

- *Set up Firebase Admin SDK and generate service account credentials.*

**4. Set up Redis (optional):**

- *Install Redis and start the Redis server (if using caching).*

- *Update your environment variables with the Redis connection details.*

**5. Set up Email service (optional):**

- *Set up a SendGrid account or SMTP email provider for notifications.*

- *Add your email API keys to the environment configuration.*

**6. Configure environment variables:**
- *Create a **.env** file and add the necessary configurations like Firebase credentials, Redis* connection, and email API keys.

---
### Running the Project
*To run the API locally, use the following command:*

```bash
go run main.go
```
*The server will start on **localhost:8080** by default.*

### Testing
*To run the unit tests:*

```bash
go test ./tests
```

### API Documentation
*The API exposes the following endpoints:*

- *POST **/auth/login**: User login with email and password.*

- *POST **/subscriptions**: Create a new subscription for a user.*

- *GET **/subscriptions/{id}**: Retrieve a user's subscription details.*

- *POST **/subscriptions/renew/{id}**: Renew a user's subscription.*

- *POST **/subscriptions/revoke/{id}**: Revoke (cancel) a user's subscription.*

*For more detailed API documentation, see API Docs.*

### Contributing
*We welcome contributions from the community! To contribute:*

- *Fork this repository.*

- *Create a new branch **(git checkout -b feature-name).***

- *Make your changes.*

- *Commit your changes **(git commit -am 'Add new feature').***

-*Push to the branch **(git push origin feature-name).***

- *Create a new Pull Request.*

#### Please ensure that your code follows the existing style and is well-tested.


### License
*This project is licensed under the **MIT License** - see the [LICENSE](https://opensource.org/license/mit) file for details.*

### Acknowledgments
- *[Firebase](https://firebase.google.com/docs/firestore) for providing Firestore.*

- *[Redis](https://redis.io/) for caching.*

- *[SendGrid](https://sendgrid.com/en-us) for email notifications.*

