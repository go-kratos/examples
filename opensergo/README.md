1. Download code: `git clone https://github.com/opensergo/opensergo-dashboard.git`
2. Modify `opensergo-dashboard-server/src/main/resources/application.yaml`, specify the mysql server address
    * Table struct: [opensergo-dashboard-server/src/main/resources/schema.sql](./opensergo-dashboard-server/src/main/resources/schema.sql)
3. Build
    * `mvn clean package -Dmaven.test.skip=true`
4. Launch
    * `cd opensergo-dashboard-server/target/; java -jar opensergo-dashboard.jar`
5. Run Service `go run main.go`

6. Visit `http://localhost:8080/`