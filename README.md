# At the beginning

## Go API with Unit Tests in Docker(14.1MB)

- Clone repo

  ```bash
  git clone https://github.com/3sky/psychic-waffle.git
  ```
- Go into `3sky` dir

  ```bash
  cd 3sky
  ```
  
- Build image

  ```bash
  docker build -t <some name> .
  ```

- Unit test run while building image

  ```console
  === RUN   TestGetUser1
  --- PASS: TestGetUser1 (0.00s)
  === RUN   TestGetUser2
  --- PASS: TestGetUser2 (0.00s)
  === RUN   TestGetUser3
  --- PASS: TestGetUser3 (0.00s)
  === RUN   TestGetUser4
  --- PASS: TestGetUser4 (0.00s)
  === RUN   TestGetUser5
  --- PASS: TestGetUser5 (0.00s)
  PASS
  ```

- Run container

  ```bash
  docker run -d -p <des port>:1323 <some name>
  ```
