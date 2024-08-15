    This project is part of learng `usegolang.com` course
    Below is the lesson plan for the same :
     The Lessons

A whole lot of them!
## Section 1: Getting Started
1. A Basic Web Application (Sample)
2. Troubleshooting and Slack
3. Packages and Imports (Sample)
4. Editors and Automatic Imports (Sample)
5. The "Hello, World" Part of our Code (Sample)
6. Web Requests (Sample)
7. HTTP Methods (Sample)
8. Our Handler Function (Sample)
9. Registering our Handler Function... (Sample)
10. Go Modules (Sample)
## Section 2: Adding New Pages
1. Dynamic Reloading (Sample)
2. Setting Header Values (Sample)
3. Creating a Contact Page (Sample)
4. Examining the http.Request Type (Sample)
5. Custom Routing (Sample)
6. URL Path vs RawPath
7. Not Found Page
8. The http.Handler Type
9. The http.HandlerFunc Type
10. Exploring Handler Conversions
11. FAQ Exercises
## Section 3: Routers and 3rd Party Libraries
1. Defining our Routing Needs
2. Using git
3. Installing Chi
4. Using Chi
5. Chi Exercises
## Section 4: Templates
1. What are Templates?
2. Why Do We Use Server Side Rendering?
3. Creating Our First Template
4. Cross Site Scripting (XSS)
5. Alternative Template Libraries
6. Contextual Encoding
7. Home Page via Template
8. Contact Page via Template
9. FAQ Page via Template
10. Template Exercises
## Section 5: Code Organization
1. Code Organization
2. MVC Overview
3. Walking Through a Web Request with MVC
4. MVC Exercises
## Section 6: Starting to Apply MVC
1. Creating the Views Package
2. fmt.Errorf
3. Validating Templates at Startup
4. Must Functions
5. Exercises
## Section 7: Enhancing our Views
1. Embedding Template Files
2. Variadic Parameters
3. Named Templates
4. Dynamic FAQ Page
5. Reusable Layouts
6. Tailwind CSS
7. Utility-first CSS
8. Adding a Navigation Bar
9. Exercises
## Section 8: The Signup Page
Creating the Signup Page
Styling the Signup Page
Intro to REST
Users Controller
Decouple with Interfaces
Parsing the Signup Form
URL Query Params
Exercises
Section 9: Databases and PostgreSQL
Intro to Databases
Intalling Postgres
Connecting to Postgres
Update: Docker Container Names
Creating SQL Tables
Postgres Data Types
Postgres Constraints
Creating a Users Table
Inserting Records
Querying Records
Filtering Queries
Updating Records
Deleting Records
Additional SQL Resources
Section 10: Using Postgres with Go
Connecting to Postgres with Go
Imports with Side Effects
Postgres Config Type
Executing SQL with Go
Inserting Records with Go
SQL Injection
Acquire a new Record's ID
Querying a Single Record
Creating Sample Orders
Querying Multiple Records
ORMs vs SQL
Exercises
Syncing the Book and Screencasts Source Code
Section 11: Securing Passwords
Steps for Securing Passwords
Third Party Authentication Options
What is a Hash Function?
Store Password Hashes, Not Encrypted or Plaintext Values
Salt Passwords
Learning bcrypt with a CLI
Hashing Passwords with bcrypt
Comparing a Password with a bcrypt Hash
Section 12: Adding Users to our App
Defining the User Model
Creating the UserService
Create User Method
Postgres Config for the Models Package
UserService in the Users Controller
Create Users on Signup
Sign In View
Authenticate Users
Process Sign In Attempts
Section 13: Remembering Users with Cookies
Stateless Servers
Creating Cookies
Viewing Cookies with Chrome
Viewing Cookies with Go
Securing Cookies from XSS
Cookie Theft
CSRF Attacks
CSRF Middleware
Providing CSRF to Templates via Data
Custom Template Functions
Adding the HTTP Request to Execute
Request Specific CSRF Template Function
Template Function Errors
Securing Cookies from Tampering
Cookie Exercises
Section 14: Sessions
Random Strings with crypto/rand
Exploring math/rand
Wrapping the crypto/rand package
Why Do We Use 32 Bytes for Session Tokens?
Defining the Sessions Table
Stubbing the SessionService
Sessions in the Users Controller
Cookie Helper Functions
Create Session Tokens
Refactor the rand Package
Hash Session Tokens
Insert Sessions into the Database
Update Existing Sessions
Query Users via Session Token
Deleting Session
Sign Out Handler
Sign Out Link
Session Exercises
Section 15: Improved SQL
SQL Relationships
Foreign Keys
On Delete Cascade
Inner Join
Left, Right, and Full Outer Join
Using Join in the SessionService
SQL Indexes
Creating PostgreSQL Indexes
On Conflict
Improved SQL Exercises
Section 16: Schema Migrations
What are Schema Migrations?
How Migration Tools Work
Installing pressly/goose
Converting to Schema Migrations
Schema Versioning Problem
Run Goose with Go
Embedding Migrations
Go Migration Files
Removing Old SQL Files
Section 17: Current User via Context
Using Context to Store Values
Improved Context Keys
Context Values with Types
Storing Users as Context Values
Reading Request Context Values
Set the User via Middleware
Requiring a User via Middleware
Accessing the Current User in Templates
Request-Scoped Values
Section 18: Sending Emails to Users
Password Reset Overview
SMTP Services
Building Emails with SMTP
Sending Emails with SMTP
Building an Email Service
EmailService.Send
Forgot Password Email
ENV Variables
Section 19: Completing the Authentication System
Password Reset DB Migration
Password Reset Service Stubs
Forgot Password HTTP Handler
Asynchronous Emails
Forgot Password HTML Template
Initializing Services with ENV Vars
Check Your Email HTML Template
Reset Password HTTP Handlers
Reset Password HTML Template
Update Password Function
Implementing PasswordReset.Create
Implementing PasswordReset.Consume
Section 20: Better Errors
Inspecting Errors
Inspecting Wrapped Errors
Designing the Alert Banner
Dynamic Alerts
Removing Alerts with JavaScript
Detecting Existing Emails
Accepting Errors in Templates
Public vs Internal Errors
Creating Public Errors
Using Public Errors
Better Error Handling Exercises
Section 21: Galleries
Galleries Overview
Gallery Model and Migration
Creating Gallery Records
Querying for Galleries by ID
Querying Galleries by UserID
Updating Gallery Records
Deleting Gallery Records
New Gallery Handler
views.Template Name Bug
New Gallery Template
Gallery Routing and CSRF Bug Fixes
Create Gallery Handler
Edit Gallery Handler
Edit Gallery Template`
Update Gallery Handler
Gallery Index Handler
Discovering and Fixing a Gallery Index Bug
Gallery Index Template Continued
Show Gallery Handler
Show Gallery Template and a Tailwind Update
Extracting Common Gallery Code
Extra Gallery Checks with Functional Options
Delete Gallery Handler
Gallery Exercises
Section 22: Images
Images Overview
Setting Up Test Images
Adding the ImagesDir to the GalleryService
Globbing Image Files
Adding Filename and GalleryID to the Image Type
Adding Images to the Show Gallery Page
Show Image Handler
Querying for a Single Image
URL Path Escaping Image Filenames
Adding Images to the Edit Gallery Page
Delete Image Form
Delete Image Service Func
Delete Image Handler
Checking for Filename Vulnerabilities
Upload Image Form
Image Upload Handler
Creating Images in the GalleryService
Detecting Content Type
Rendering Content Type Errors
Section 23: Preparing for Production
Loading All Config via ENV
Docker Compose Overrides
Building Tailwind Locally
Tailwind Via Docker
Serving Static Assets
Making main Easier to Test
Running our Go Server via Docker
Multi-Stage Docker Builds
Tailwind Production Build
Caddy Server via Docker
Section 24: Deploying
Creating a Digital Ocean Droplet
Setting up DNS
Installing Git on the Server
Setting Up a Bare Git Repo
Setting Up a Local Git Repo
Checking Out Our Code on the Server
Email Sending Server Setup
Production .env File
Install Docker in Prod
Production Caddyfile
Production Data Directories
Running Our App in Prod
Post-receive Deploy Updates
Deploy via Git
Logging Services
Bonus: OAuth
Intro to OAuth
OAuth Example Code
Dropbox App Setup
Offline OAuth Demo
OAuth Tokens
Online vs Offline Access Types
Redirect URIs
OAuth Connect HTTP Handler
Determine Redirect URI Host
OAuth Routes and Config Setup
OAuth Callback Handler
Testing OAuth with API Calls
Bonus: Dropbox Chooser
Dropbox Chooser Overview
Embedding the Chooser
Images via Dropbox Form
Chooser Success Function
Images via URL Handler
Downloading Images
Create Images Without Seek
Concurrent Downloads
Using errgroup
Page Specific JS 

# lenslocked
**Need clear picture about ServeMux , defaultServeMux and handler func while**
creating golang server:
https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go
//Try to go through again once after finishing the usegolang course

## Dynamic Loading
    With help of modd.

## Response Writer
    `Header().Set()`
## Request Method
## Routing
    Determining  Which page to show to user
## wr**
     make w & r for handler
## 3rd party router**: 
- Chi
- gorilla mux

## Explore how can u make you own router?


## Template
- text/

## xss Attack :
https://owasp.org/www-community/attacks/xss/
In this attack hacker inject the malicious script in the application can easily
access the session token and cookies and even delete the user account .

## Contextual-encoding 
https://www.calhoun.io/intro-to-templates-p1-contextual-encoding/
refer : branch contextual-encoding

