# Changelog


## [0.2.0-chapter2](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/compare/v0.1.0-chapter1..v0.2.0-chapter2) - 2024-10-01




### Tests

- *(parser)* Add generic functions for testing expressions - ([67c759b](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/67c759b26cc6b254315c32d67c1b9b5da377151f))

### ✨ Features

- *(repl)* More monkey - ([ca0e661](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/ca0e6614a979c4ae6b405f4eb71596132bdcce9e))
- *(repl)* Utilize parser in loop - ([92e2b5a](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/92e2b5ae1dd2512206f9c54f26f8b78d8741a24d))
- *(parser)* Finalize parser - ([ceeb787](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/ceeb7874103c5280b86123613768302c1fce4e71))
- *(parser)* Parse call expressions - ([3e1721a](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/3e1721afaab4cfe4f4dbe3db45e6c87da3ce02cc))
- *(parser)* Parse functions - ([794eebf](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/794eebfeef305880377fccdc5904dae3997e2345))
- *(parser)* Parse else part of if expressions - ([728965c](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/728965c23729c0e234b786bb0a95a66b626d02cd))
- *(parser)* Parse simple if expressions - ([accc2c2](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/accc2c28a93e8049cbeee4482a2b28b5661adc98))
- *(parser)* Parse grouped expressions - ([704c3ad](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/704c3ade5ea279524dcef086b27f5d43803823a1))
- *(parser)* Parse boolean expressions - ([d2ad954](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/d2ad954c0ab3af72543f2ff9a71570e1568062a5))
- *(parser)* Easy functions for debugging of pratt parser - ([35ae652](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/35ae652204737bebd4ec7352b479d20d0484a643))
- *(parser)* Parse infix expressions - ([1bc58a7](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/1bc58a74561dad61708391f63c31a9dc8571526d))
- *(parser)* Map tokentypes to precedence constants - ([5e638d3](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/5e638d3937f6ec0aa3cc163324d79c0b6e29d45d))
- *(parser)* Parse prefix expressios (! and -) - ([8369cb8](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/8369cb85a6b92dd54925f0ead696e22cdaee139d))
- *(parser)* Add error handling if there is no parser for an expression - ([36d6da5](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/36d6da5b9de76c42f71b93c6087096cc3aedc849))
- *(parser)* Parse integer expressions - ([cc5e389](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/cc5e389afc9d65b809bdf57893b5370f9b25f281))
- *(parser)* Parse singule identifier as expression - ([58d2d9f](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/58d2d9f56d574c60f3e0cb19f925e186bca6d194))
- *(parser)* Add debugging methods for identifiers and expressions - ([c2b3b1a](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/c2b3b1a652c58d80722dc7a61794d0e91d400352))
- *(parser)* Parse return statements without expressions - ([183a670](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/183a670e73172928f27d421832f1b6a22bdefd36))
- *(parser)* Add basic error handling - ([ee643ce](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/ee643ce79d650583f4dc2b739497ba1cda71a646))
- *(parser)* (unfinished) parse simple let statemets - ([e937e52](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/e937e52387aa090790f21ab967fdbf8a59abcc0c))
- *(parser)* Add basic parser structure - ([8edc37b](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/8edc37b71a437c9ceb9a0caa190effc3298adb9c))
- *(ast)* Imnplement barebone ast - ([05ef76e](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/05ef76ebb704202eb8efb49ec19fd581d711f0bc))
- *(repl)* Add a read eval print loop - ([69c9895](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/69c9895275dbe5d6a1716b61b245b833606f0c4a))
- *(lexer)* Lex == and != - ([984ac2d](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/984ac2d03b64c4f4cbeb53936bb8b2a5dbbc73a0))
- *(token)* Add true, if, else and return - ([c2d0fba](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/c2d0fba9c2435cbbcbcd52086340bcd744853d7f))
- *(lexer)* Lex -, /, *, < and > - ([87e0758](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/87e075812b4e7eeae00c35d374daf697b2e6a9e0))
- *(lexer)* Lex keywords and identifiers - ([b13b15a](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/b13b15a93e5a89b2a027d51fa129183c408ee5bf))
- *(lexer)* Skip whitespaces - ([0933501](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/0933501a957ec24f15592dff24f3b5714a11bd0b))
- *(lexer)* Implement basic lexer - ([e68a6c4](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/e68a6c4ba6d8ef808dd330c7bc48588cdc9e9f75))
- *(token)* Define token literals - ([d6f8a1d](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/d6f8a1dbd1fd05cd471dac1be3241d7ea76290f5))

### 🎉 Initial Commit

- *(go)* Add go related boilerplate - ([9a80c4c](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/9a80c4c694b9fceabce968252edba2e2c8cbc0f2))
- *(git-cliff)* Create git-cliff template - ([1e79f0c](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/1e79f0c136960a096aadf05032d9ff0b91e365f2))

### 🐛 Bug Fixes

- *(cliff)* Correct repo name in cliff config - ([3605470](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/3605470de703b9db79857b32bc9ed20cb560b8b7))
- *(ast)* Fix typo in ExpressionStatement interface implementation - ([84b71e9](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/84b71e970a19e87e8e2124fd69b2d81cab0ac914))
- *(ast)* Add missing return type for function implented from interface - ([ede24db](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/ede24db0f1ac887cdcdbe3b13bf01c297a284519))
- *(token)* Remove typo in TokenType struct - ([cab737d](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/cab737d02bf177cda7228df9796e2669d2f1ba0d))

### 🔧Chores

- *(ci)* Change auto-generated releases to be pre-release - ([79fc75c](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/79fc75c0db94c4ba2d2d9d7109d1955434f588f0))
- *(repo)* Update README with source code examples from official website - ([faddf2b](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/faddf2bb5501c89747a60401a3a1872bccd1893c))
- *(repo)* Add README - ([7afa27c](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/7afa27cf11492ccff727ab985e13e7b0ae68f5a6))
- *(release)* Prepare for v0.1.0-chapter1 - ([396a913](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/396a913974a1d5dfab93af6b5f7c76e71270dae1))

### 🚜 Refactoring

- *(core)* Rename vars - ([96ebb88](https://github.com/Flokkq/https://github.com/Flokkq/monkey-interpreter/commit/96ebb880993cb5ca3216e9e1edfbe25a60fd20e7))
<!-- generated by git-cliff -->
