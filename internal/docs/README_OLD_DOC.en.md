# gormcngen: Provides a Columns() Function to Retrieve Column Names for GORM Models

Before using this tool, please familiarize yourself with the project: [gormcnm](https://github.com/yyle88/gormcnm).

Assuming you understand the purpose of `gormcnm`, we will now introduce how to use `gormcngen`.

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
- Before running the test, make sure the target code file (`gormcnm.gen.go`) exists. If not, please manually create an empty file to avoid overwriting other files.
- After installing dependencies, you can run the test to generate the code:
   ```bash
   go get github.com/yyle88/gormcngen
   ```

---

## Example Reference

For more usage examples, refer to the [examples package](../../internal/examples). This includes multiple ways to generate models, as well as demonstrations of merging the generated code with the model code into a single file.  
However, it is recommended to store the generated code separately for easier cleanup and regeneration, rather than mixing it with model definition files.

---

## Easter Eggs

### Using Chinese Encoding

Example: [Chinese Encoding Example](../../internal/examples/example4/example4usage/example4usage_test.go) shows how to use Chinese encoding. This is a simple exploration.  
While many believe learning English is important, for developers, quickly completing business tasks and achieving financial freedom may be more practical.

### Using Native Language Encoding

You can also write code in other languages.  
Using your native language not only helps improve efficiency but also makes complex business logic easier to implement.  
For example, when a single developer writes code in English, the project code limit is about 50,000 lines. Using your native language could increase this limit several times, making it ideal for development during fragmented time.

---

## Additional Usage

When integrating with `gorm`, refer to the [gormcnm project](https://github.com/yyle88/gormcnm). Although the project documentation may not be detailed enough, you can gradually master it through practice.

---

## Other Notes

After open-sourcing this tool, I can more easily use it in company projects, and the benefits have far exceeded expectations.

As for why [gormcnm](https://github.com/yyle88/gormcnm) and [gormcngen](https://github.com/yyle88/gormcngen) are not merged into a single project, the reason is that the tool and generation packages were initially designed to be separate. This structure was retained when open-sourced. Of course, these two tools are not required to be used together, so it is most appropriate to keep them independent.

---
