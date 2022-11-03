import { Config } from "../../config"
import { Sample } from "../../sample"

Sample.run("builder_from_hex", async (config: Config) => {
    console.log('start from_hex', config)
    setTimeout(() => {
        console.log('end from_hex')
    }, 1000)
})