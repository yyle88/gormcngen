# gormcngen: Provides a Columns() Function to Retrieve Column Names from GORM Models

Before using this tool, please familiarize yourself with the project: [gormcnm](https://github.com/yyle88/gormcnm).

Assuming you understand the purpose of `gormcnm`, we now introduce how to use `gormcngen`.

## Quick Start

For the guide, refer to the [demo package](../../internal/demos/demo1).

1. **View Example Model**  
   Example model code: [example.go](../../internal/demos/demo1/demo1models/example.go)

2. **Understand the Code Generation Process**  
   Test generated code: [gormcnm.gen_test.go](../../internal/demos/demo1/demo1models/gormcnm.gen_test.go)

3. **View Generated Results**  
   Generated intermediate code: [gormcnm.gen.go](../../internal/demos/demo1/demo1models/gormcnm.gen.go)

4. **Run Business Logic**  
   Main program logic: [main.go](../../internal/demos/demo1/main.go)

### Simplest Usage

Simply copy the `gormcnm.gen_test.go` file from the [models package](../../internal/demos/demo1/demo1models) to your project's `model/models` package. Then, write the models you need to generate into the test file.

**Notes:**
- Modify the package name in the test file according to your project needs, for example, replace `models` with your model package name.
- Before running the test, make sure the target code file (`gormcnm.gen.go`) exists. If not, create a blank file to avoid overwriting different files.
- Once dependencies are installed, you can run the test to generate the code:
   ```bash
   go get github.com/yyle88/gormcngen
   ```

---

## Example Reference

For more usage examples, refer to the [examples package](../../internal/examples). This includes multiple ways to generate models, as well as demonstrations of merging the generated code with the model code into a single file.  
We recommend storing the generated code in separate files to make cleanup and regeneration smooth, instead of mixing it with definition files.

---

## Bonus Features

### Using Chinese Encoding

Example: [Chinese Encoding Example](../../internal/examples/example4/example4usage/example4usage_test.go) shows how to use Chinese encoding. This is a simple exploration.  
While some believe learning English is important, completing business tasks fast and achieving business success might be more strategic.

### Using Native Language Encoding

You can also write code in different languages.  
Using native language not just helps improve performance but also makes complex business logic smooth to implement.  
As an example, when a single person writes code in English, the project code limit is about 50,000 lines. Using native language could increase this limit multiple times, making it suitable during fragmented time.

---

## Extra Usage

When integrating with `gorm`, refer to the [gormcnm project](https://github.com/yyle88/gormcnm). Although the project documentation may not be detailed enough, you can gradually master it through practice.

---

## Extra Notes

Once this package was open-sourced, I can use it in business projects with ease, and the benefits have exceeded expectations.

As for why [gormcnm](https://github.com/yyle88/gormcnm) and [gormcngen](https://github.com/yyle88/gormcngen) are not merged into a single project, the reason is that the tool and generation packages were initially designed to be separate. This structure was retained when open-sourced. Of course, these two tools are not required to be used together, so it is most appropriate to keep them independent.

---
