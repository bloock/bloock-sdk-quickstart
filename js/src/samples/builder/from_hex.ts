import { Config } from "../../utils/config"
import { Sample } from "../../utils/sample"

Sample.run("builder_from_hex", async (config: Config) => {
    console.log('start from_hex', config)
    setTimeout(() => {
        console.log('end from_hex')
    }, 2000)
})