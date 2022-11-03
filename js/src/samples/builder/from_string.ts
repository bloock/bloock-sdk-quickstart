import { Config } from "../../utils/config"
import { Sample } from "../../utils/sample"

Sample.run("builder_from_string", async (config: Config) => {
    console.log('start from_string', config)
    setTimeout(() => {
        console.log('end from_string')
    }, 2000)
})