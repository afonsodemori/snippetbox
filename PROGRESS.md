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
   - [x] Displaying dynamic data - [74a8fc1]
   - [x] Template actions and functions - [7c3dd38]
   - [x] Caching templates - [38dde2b]
   - [x] Catching runtime errors - [8137361]
   - [x] Common dynamic data - [a8641ec]
   - [x] Custom template functions - [c2816a3]
6. Middleware
   - [x] How middleware works - <small>`theory`</small>
   - [x] Setting common headers - [2111b3f]
   - [x] Request logging - [a347b03]
   - [x] Panic recovery - [9ac90fe]
   - [x] Composable middleware chains - [04ded9b]
7. Processing forms
   - [x] Setting up an HTML form - [6aee7f4]
   - [x] Parsing form data - [d1f08fa]
   - [x] Validating form data - [8291b7c]
   - [x] Displaying errors and repopulating fields - [310dbd8]
   - [x] Creating validation helpers - [1a7f28f]
   - [x] Automatic form parsing - [a1ccf7c]
8. Stateful HTTP
   - [x] Choosing a session manager - [3e2e33d]
   - [x] Setting up the session manager - [3f68ad9]
   - [x] Working with session data - [3a56e7e]
9. Server and security improvements
   - [x] The http.Server struct - [aeb991b]
   - [x] The server error log - [0c1a69c]
   - [x] Generating a self-signed TLS certificate - [3c2ee35]
   - [x] Running an HTTPS server - [4112667]
   - [x] Configuring HTTPS settings - [f224490]
   - [x] Connection timeouts - [d442270]
10. User authentication
    - [x] Routes setup - [1c9ffc5]
    - [x] Creating a users model - [195a6c9]
    - [x] User signup and password encryption - [ad2a2ee]
    - [x] User login - [1ab0aca]
    - [x] User logout - [cc6ed70]
    - [x] User authorization - [b7917e5]
    - [x] CSRF protection - [ddbc5d4]
11. Using request context
    - [x] How request context works - <small>`theory`</small>
    - [x] Request context for authentication/authorization - [27daeaf]
12. File embedding
    - [x] Embedding static files - [f6848b5]
    - [x] Embedding HTML templates - [6139ba7]
13. Testing
    - [x] Unit testing and sub-tests - [208643b]
    - [x] Assertion helpers - [e64c8a8]
    - [x] Testing HTTP handlers and middleware - [e8099db]
    - [x] End-to-end testing - [686232c]
    - [x] Customizing how tests run - <small>`theory`</small>
    - [x] Mocking dependencies - [3b28076]
    - [x] Testing HTML forms - [16b932f]
    - [x] Integration testing - [5a55461]
    - [x] Profiling test coverage - <small>`theory`</small>
14. Conclusion
15. Further reading and useful links
16. Guided exercises
    - [x] Add an 'About' page to the application - [c80263a]
    - [ ] Add a debug mode
    - [ ] Test the snippetCreate handler
    - [x] Add an 'Account' page to the application - [ee895c1]
    - [x] Redirect user appropriately after login - [5ee4ea0]
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
[74a8fc1]: https://github.com/afonsodemori/snippetbox/commit/74a8fc1706f34cbc591d0f25f3a043ea926d29f2
[7c3dd38]: https://github.com/afonsodemori/snippetbox/commit/7c3dd38027eaca29fa516ea8e401fdf8ff6da272
[38dde2b]: https://github.com/afonsodemori/snippetbox/commit/38dde2bda04410bc4e2d8ae2e8590248af50e62e
[8137361]: https://github.com/afonsodemori/snippetbox/commit/8137361c80e65cf226d2e3c02025e4dfd14b3d42
[a8641ec]: https://github.com/afonsodemori/snippetbox/commit/a8641ecbdc31d2487e9f85072d03ff72ad0b008a
[c2816a3]: https://github.com/afonsodemori/snippetbox/commit/c2816a3c7cd5dd71d9692839c7dc223c832ba654
[2111b3f]: https://github.com/afonsodemori/snippetbox/commit/2111b3f90d98d9e65130dacdc9d4539c2a6e4d65
[a347b03]: https://github.com/afonsodemori/snippetbox/commit/a347b03442e48a6ca65dbee038b961392f930505
[9ac90fe]: https://github.com/afonsodemori/snippetbox/commit/9ac90feffb63e434e21b0e8352f568d8d0c342ec
[04ded9b]: https://github.com/afonsodemori/snippetbox/commit/04ded9b9eb8e790efb9989c8cabff445c960ddc5
[6aee7f4]: https://github.com/afonsodemori/snippetbox/commit/6aee7f45f12709f91aafacf99a37c460230e604a
[d1f08fa]: https://github.com/afonsodemori/snippetbox/commit/d1f08fa94b27cc28a0cc3673315189308aab419e
[8291b7c]: https://github.com/afonsodemori/snippetbox/commit/8291b7c182ff1f7a3c9c47c0ce69844f4bba741f
[310dbd8]: https://github.com/afonsodemori/snippetbox/commit/310dbd86926813f33a0b3de95bbc5a4b98a0e7cc
[1a7f28f]: https://github.com/afonsodemori/snippetbox/commit/1a7f28ff9ec3b721d794c568e4000d44398586ad
[a1ccf7c]: https://github.com/afonsodemori/snippetbox/commit/a1ccf7c79e2ab0d4543e599b2223e37b29e0ad5f
[3e2e33d]: https://github.com/afonsodemori/snippetbox/commit/3e2e33dd3bd80ce3febd1469eea3346e81c62231
[3f68ad9]: https://github.com/afonsodemori/snippetbox/commit/3f68ad993591f37b843c04779f7a485037b8692d
[3a56e7e]: https://github.com/afonsodemori/snippetbox/commit/3a56e7ef714d93e097cc92c3a1ed30d1bfda8656
[aeb991b]: https://github.com/afonsodemori/snippetbox/commit/aeb991b74e479b939f8638917df54c6783b31bd8
[0c1a69c]: https://github.com/afonsodemori/snippetbox/commit/0c1a69c3c17129c1620e27bbb8d23909f74c0fbc
[3c2ee35]: https://github.com/afonsodemori/snippetbox/commit/3c2ee354dc6eb4bfe142239f62a832b2e1a7569c
[4112667]: https://github.com/afonsodemori/snippetbox/commit/4112667d19cc70bdceab644c017feb29fe70574f
[f224490]: https://github.com/afonsodemori/snippetbox/commit/f2244905d005cc557f9c27c0a57c247e97d84e25
[d442270]: https://github.com/afonsodemori/snippetbox/commit/d442270faba861d554c0dc6b034a25bd96208182
[1c9ffc5]: https://github.com/afonsodemori/snippetbox/commit/1c9ffc5f08ea0497cf955dd90c8313a739fe8a76
[195a6c9]: https://github.com/afonsodemori/snippetbox/commit/195a6c9c8b29cd54ea9300fe2e89cdacb6464a93
[ad2a2ee]: https://github.com/afonsodemori/snippetbox/commit/ad2a2ee210704da8fdab4df6cca39d3fb9730cee
[1ab0aca]: https://github.com/afonsodemori/snippetbox/commit/1ab0aca1077f2afccdd90d82e2d475c51fc57738
[cc6ed70]: https://github.com/afonsodemori/snippetbox/commit/cc6ed7090101e2283d48f094ae3fee5f00d948e6
[b7917e5]: https://github.com/afonsodemori/snippetbox/commit/b7917e5edb8eec8548d64b69179c5d0b59c6c10f
[ddbc5d4]: https://github.com/afonsodemori/snippetbox/commit/ddbc5d4dc0eba9843bf4f776085ec73e0932fd8c
[27daeaf]: https://github.com/afonsodemori/snippetbox/commit/27daeaf4794b6ecb38e2f3a959acf91b937746ca
[f6848b5]: https://github.com/afonsodemori/snippetbox/commit/f6848b5660fe0ce06e725aa70778cceb312bc508
[6139ba7]: https://github.com/afonsodemori/snippetbox/commit/6139ba70ddb28ebbe5c2f91b78d4f14af2e9984b
[208643b]: https://github.com/afonsodemori/snippetbox/commit/208643bb8165bd18274a8d65198a8cde7ca19336
[e64c8a8]: https://github.com/afonsodemori/snippetbox/commit/e64c8a89474e7280c9c5950f471bddcf6246518d
[e8099db]: https://github.com/afonsodemori/snippetbox/commit/e8099dbe5f9a172dad372889b587d1254344c20e
[686232c]: https://github.com/afonsodemori/snippetbox/commit/686232c958a1a4ddec3720149b700f6818b04d3f
[3b28076]: https://github.com/afonsodemori/snippetbox/commit/3b28076ef96be46faa15470ba9ccf7f8cf92d5b2
[16b932f]: https://github.com/afonsodemori/snippetbox/commit/16b932f4b303b4f7440d6c840996b92cd210c881
[5a55461]: https://github.com/afonsodemori/snippetbox/commit/5a55461eb3685da94e563b90c514b11eaf01d2ee
[c80263a]: https://github.com/afonsodemori/snippetbox/commit/c80263a9a7901b6c2e1171e39621c5715550d655
[ee895c1]: https://github.com/afonsodemori/snippetbox/commit/ee895c1c2d11fec3b2beec6887c67bff82c81dcb
[5ee4ea0]: https://github.com/afonsodemori/snippetbox/commit/5ee4ea0acb88084b8b5bbb23f0000e87fca3f687
