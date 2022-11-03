import { Config } from "../../config"
import { Sample } from "../../sample"

Sample.run("builder_from_string", async (config: Config) => {
    console.log('start from_string', config)
    return Promise.reject("an error")
    setTimeout(() => {
        console.log('end from_string')
    }, 1000)
})