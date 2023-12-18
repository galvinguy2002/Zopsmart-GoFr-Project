# Zopsmart GoFr Project
 A Simple GoFr Mini Project built by a Go Getter !! ![1_dX7_Le4fztcY0QGbuk6KWA](https://github.com/galvinguy2002/Zopsmart-GoFr-Project/assets/119154626/8f2ec3d7-e5b2-4bfd-a119-ea370324247a)


 ## Net Requirements:
Build a simple HTTP (REST) API using GoFr--
### 1. CRUD Operations: Build APIs for create, read, update and delete operations for all entities
  - [x] Create :white_check_mark:
        
  - [x] Read :white_check_mark:
        
  - [x] Update/Delete :white_check_mark:
  

### 2. DB Integration: The API should have integration with database for persistence of data. Any SQL or NoSQL DB, which has a freely available docker image, may be used
  - [x] Database Creation :white_check_mark:
        
  - [x] Integration using Gofr :white_check_mark:
  

### 3. Unit Tests: A minimum of 60% unit test coverage is desired
- [x] All CRUD op represent approximate unit test coverage :white_check_mark:

- [x] Try to incorporate a real-world scenario while implementing the project. :white_check_mark:
      
    - [x] The real world scenario involves Stock Management API where user can retrieve details in relation with STOCKS. :white_check_mark:
## Extra Points

- [ ] Unit Test Coverage > 90%
      
- [x] Postman collection for trying out the APIs(partially done) :heavy_exclamation_mark::heavy_check_mark:
      
- [x] Sequence diagrams, UML diagrams :white_check_mark:

- [x] Subtle Front-End components provided for easy understanding :white_check_mark:
 <br>
[[ The following sequence diagram represents the outline of the API as complexity of the entire application and the interactions with the gofr package and the database is not possible here. ]] <br>
![SEQ diag](https://github.com/galvinguy2002/Zopsmart-GoFr-Project/blob/main/img/seq.png)

 <br>
--> Additional UML Diagram for visual understanding : <br>
![UML diag](https://github.com/galvinguy2002/Zopsmart-GoFr-Project/blob/main/img/UML.png)

### How to Run the Stock API:

**Prerequisites:**
- Ensure you have Go installed on your system. If not, you can download and install it.
FOLLOW IN ORDER
1. **Clone the Repo**
   - Clone the repository containing the Go code for the Movie Management API.

2. **Install Dependencies:**
   - If your project has any external dependencies, use the following command to install them.
   ```bash
   go get ./...
   -->gofr
   -->npm packages (if req)
   -->go mod tidy/go mod init for tracking
   ```
   
3. **Build and Run:**
   - Build and run the `main.go` file.
  go run main.go
   ```
   
   This will start the API server, and it will be accessible at `http://localhost:8000`.


6. **Testing with Postman:**
   - Install [Postman](https://www.postman.com/) to test the API endpoints.
   - Use Postman to send HTTP requests to the API at `http://localhost:8000/stock` for testing CRUD operations.

----------------------------------------------------------------------------------------------------------------------------------------------
### Prerequisites For FrontEnd
Make sure you have the following installed on your machine:
- Node.js & npm - [Download & Install Node.js](https://nodejs.org/)
- Git - [Download & Install Git](https://git-scm.com/)
### Installation
1. Clone the repository:

    ```bash
    git clone https://github.com/galvinguy2002/Zopsmart-GoFr-Project
    ```

2. Navigate to the project directory:

    ```bash
    cd stocks-management-app
    ```
3. Install dependencies:

    ```bash
    npm install
    ```
### Running the Application
To stat the React development server, use the following command:

```bash
npm start
 ```

## 🎉 Thanks

Built with : <br>
Golang & GoFr <br>
HAPPY CODING :) <br>

WITH REGARDS TO THE ZOPSMART TEAM <br>
Harshvardhan Saxena <br>
20104064 <br>
harshsaxena5434@gmail.com <br>
Jaypee Institute of Information Technology(JIIT), Noida
