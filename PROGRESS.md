# Progress Map

Track of my journey through [_Let's Go_](https://lets-go.alexedwards.net/). Completed sections link directly to the corresponding commit implementation.

1. Introduction
   - [x] Prerequisites - <small>`theory`</small>
2. Foundations
   - [x] Project setup and creating a module - [4f4de3b]
   - [x] Web application basics - [b25622a]
   - [x] Routing requests - [a98fc50]
   - [x] Wildcard route patterns - [47ec3d3]
   - [x] Method-based routing - [78adfab]
   - [x] Customizing responses - [17201a5]
   - [x] Project structure and organization - [c6fddaa]
   - [x] HTML templating and inheritance - [c86c743]
   - [x] Serving static files - [529f25c]
   - [x] The http.Handler interface - <small>`theory`</small>
3. Configuration and error handling
   - [x] Managing configuration settings - [4e77db3]
   - [x] Structured logging - [44ae9c6]
   - [x] Dependency injection - [b5ccf9b]
   - [x] Centralized error handling - [a2d62b4]
   - [x] Isolating the application routes - [ea101e2]
4. Database-driven responses
   - [x] Setting up MySQL - [9da4a68]
   - [x] Installing a database driver - [979bd20]
   - [x] Modules and reproducible builds - <small>`theory`</small>
   - [x] Creating a database connection pool - [9283730]
   - [x] Designing a database model - [dcf2794]
   - [x] Executing SQL statements - [48dd337]
   - [x] Single-record SQL queries - [c6b92c9]
   - [x] Multiple-record SQL queries - [6279e08]
   - [x] Transactions and other details - <small>`theory`</small>
5. Dynamic HTML templates
   - [ ] Displaying dynamic data
   - [ ] Template actions and functions
   - [ ] Caching templates
   - [ ] Catching runtime errors
   - [ ] Common dynamic data
   - [ ] Custom template functions
6. Middleware
   - [ ] How middleware works
   - [ ] Setting common headers
   - [ ] Request logging
   - [ ] Panic recovery
   - [ ] Composable middleware chains
7. Processing forms
   - [ ] Setting up an HTML form
   - [ ] Parsing form data
   - [ ] Validating form data
   - [ ] Displaying errors and repopulating fields
   - [ ] Creating validation helpers
   - [ ] Automatic form parsing
8. Stateful HTTP
   - [ ] Choosing a session manager
   - [ ] Setting up the session manager
   - [ ] Working with session data
9. Server and security improvements
   - [ ] The http.Server struct
   - [ ] The server error log
   - [ ] Generating a self-signed TLS certificate
   - [ ] Running an HTTPS server
   - [ ] Configuring HTTPS settings
   - [ ] Connection timeouts
10. User authentication
    - [ ] Routes setup
    - [ ] Creating a users model
    - [ ] User signup and password encryption
    - [ ] User login
    - [ ] User logout
    - [ ] User authorization
    - [ ] CSRF protection
11. Using request context
    - [ ] How request context works
    - [ ] Request context for authentication/authorization
12. File embedding
    - [ ] Embedding static files
    - [ ] Embedding HTML templates
13. Testing
    - [ ] Unit testing and sub-tests
    - [ ] Assertion helpers
    - [ ] Testing HTTP handlers and middleware
    - [ ] End-to-end testing
    - [ ] Customizing how tests run
    - [ ] Mocking dependencies
    - [ ] Testing HTML forms
    - [ ] Integration testing
    - [ ] Profiling test coverage
14. Conclusion
15. Further reading and useful links
16. Guided exercises
    - [ ] Add an 'About' page to the application
    - [ ] Add a debug mode
    - [ ] Test the snippetCreate handler
    - [ ] Add an 'Account' page to the application
    - [ ] Redirect user appropriately after login
    - [ ] Implement a 'Change Password' feature

[4f4de3b]: https://github.com/afonsodemori/snippetbox/commit/4f4de3b1c43bffff43bb315a07c860e2283ad1aa
[b25622a]: https://github.com/afonsodemori/snippetbox/commit/b25622ad04f52213091253276c3e6f2bb1ffc20c
[a98fc50]: https://github.com/afonsodemori/snippetbox/commit/a98fc50bb1634ffd526be3c2a00abfb7e65a3109
[47ec3d3]: https://github.com/afonsodemori/snippetbox/commit/47ec3d326c34e824fff806476dcf3b8b4ac48ce0
[78adfab]: https://github.com/afonsodemori/snippetbox/commit/78adfabd866cb56a35fd70c54e3464b4945dbd2e
[17201a5]: https://github.com/afonsodemori/snippetbox/commit/17201a5ee32b064f20eec70195f19b028d2f8125
[c6fddaa]: https://github.com/afonsodemori/snippetbox/commit/c6fddaa95ef63a46d876a747d4af28e80dc0ad32
[c86c743]: https://github.com/afonsodemori/snippetbox/commit/c86c74357beb2a8da8cf7801a84fdb4cf3c42214
[529f25c]: https://github.com/afonsodemori/snippetbox/commit/529f25cf65660a9784c26cd583bc082d27cb2fc4
[4e77db3]: https://github.com/afonsodemori/snippetbox/commit/4e77db32e2969d416c9ad70633354af3d7b2c2b7
[44ae9c6]: https://github.com/afonsodemori/snippetbox/commit/44ae9c6a1c79634b183a0712051f8bae15f6c982
[b5ccf9b]: https://github.com/afonsodemori/snippetbox/commit/b5ccf9b83065f408c53943ff1bd9c9432d2f7e1d
[a2d62b4]: https://github.com/afonsodemori/snippetbox/commit/a2d62b4d35d2f97ec67604dff356a7b8d75ff108
[ea101e2]: https://github.com/afonsodemori/snippetbox/commit/ea101e2a98f84d25da91bc75695e30186e0d8852
[9da4a68]: https://github.com/afonsodemori/snippetbox/commit/9da4a683333f281f4f803490eb4c491d12f8cb76
[979bd20]: https://github.com/afonsodemori/snippetbox/commit/979bd20d8531ff363ddec4029cd02fc644214d43
[9283730]: https://github.com/afonsodemori/snippetbox/commit/9283730be7e13be73ef5ccec3c1b21e317b82c49
[dcf2794]: https://github.com/afonsodemori/snippetbox/commit/dcf2794a99ff6da35eeac26b9dd40e650ee3944a
[48dd337]: https://github.com/afonsodemori/snippetbox/commit/48dd337eaac29f7f285f5b0e036c4f17a0c92cb8
[c6b92c9]: https://github.com/afonsodemori/snippetbox/commit/c6b92c9bee3fecd1a738648b744187da2392e818
[6279e08]: https://github.com/afonsodemori/snippetbox/commit/6279e083a59a2dfcb9b282c17b88a17c4fb8876a
