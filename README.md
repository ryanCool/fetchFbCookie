
# Go Selenium WebDriver

This Go project demonstrates how to set up a Selenium WebDriver server, use it to navigate to a webpage, and retrieve information about the webpage's cookies and performance logs.

## Prerequisites

1. Install [Go](https://golang.org/doc/install)
2. Install [Selenium WebDriver](https://www.selenium.dev/documentation/en/selenium_installation/installing_webdriver_binaries/)
3. Download [ChromeDriver](https://chromedriver.chromium.org/downloads) compatible with the version of Chrome installed on your machine

## Usage

1. Clone this repository
2. Navigate to the cloned repository
3. Update the `seleniumPath` constant in `main.go` to the path where you downloaded ChromeDriver
4. Run the Go program using the command: `go run main.go`
5. The program will launch a Chrome browser instance, navigate to a specific Facebook page, and print out the webpage's cookies and performance logs

## Code Structure

`main.go`:
- Sets up a Selenium service using ChromeDriver
- Creates a new remote WebDriver instance
- Navigates to a Facebook page
- Prints out the webpage's cookies and performance logs

## Built With

- [Go](https://golang.org/)
- [Selenium WebDriver](https://www.selenium.dev/)

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
