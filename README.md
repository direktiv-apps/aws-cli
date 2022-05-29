
# aws-cli 1.0

Execute AWS CLI commands from Direktiv.

---
- #### Categories: cloud, aws
- #### Image: gcr.io/direktiv/apps/aws-cli 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/aws-cli/issues
- #### URL: https://github.com/direktiv-apps/aws-cli
- #### Maintainer: [direktiv.io](https://www.direktiv.io)
---

## About aws-cli

This service excutes AWS CLI commands. All commands are getting executed in the specified region and return their results as JSON.

### Example(s)
  #### Function Configuration
  ```yaml
  functions:
  - id: aws-cli
    image: gcr.io/direktiv/apps/aws-cli:1.0
    type: knative-workflow
  ```
   #### Basic
   ```yaml
   - id: req
     type: action
     action:
       function: aws-cli
       secrets: ["awsacess", "awssecret"]
       input:
        access-key: jq(.secrets.awsacess)
        secret-key: jq(.secrets.awssecret)
        region: eu-central-1
        commands:
        - command: aws ec2 describe-instances
          print: false
        - command: aws ecr get-login-password
   ```

### Request

The request body includes a list of AWS CLI commands.

#### Request Attributes
[PostParamsBody](#post-params-body)

### Response
  AWS CLI response.
#### Reponse Types
    
  

[PostOKBody](#post-o-k-body)
#### Example Reponses
    
```json
{
  "aws": [
    {
      "result": "VTQ3U....c2ZaN0FJaldjVnkra2tKV==",
      "success": true
    },
    {
      "result": "exit status 254",
      "success": false
    }
  ]
}
```

### Errors
| Type | Description
|------|---------|
| io.direktiv.command.error | Command execution failed |
| io.direktiv.output.error | Template error for output generation of the service |
| io.direktiv.ri.error | Can not create information object from request |


### Types
#### <span id="post-o-k-body"></span> postOKBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| aws | [][PostOKBodyAwsItems](#post-o-k-body-aws-items)| `[]*PostOKBodyAwsItems` |  | |  |  |


#### <span id="post-o-k-body-aws-items"></span> postOKBodyAwsItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| result | [interface{}](#interface)| `interface{}` | ✓ | |  |  |
| success | boolean| `bool` | ✓ | |  |  |


#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| access-key | string| `string` | ✓ | | AWS access key. | `ABCABCABCDABCABCABCD` |
| commands | [][PostParamsBodyCommandsItems](#post-params-body-commands-items)| `[]*PostParamsBodyCommandsItems` |  | | Array of commands. |  |
| files | [][DirektivFile](#direktiv-file)| `[]apps.DirektivFile` |  | | File to create before running commands. |  |
| region | string| `string` |  | `"us-east-1"`| Region the commands should be executed in. | `eu-central-1` |
| secret-key | string| `string` | ✓ | | AWS secret key. | `Abcd45sa01234+ThIsIsSuPeRsEcReT` |


#### <span id="post-params-body-commands-items"></span> postParamsBodyCommandsItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| command | string| `string` |  | | Command to run | `aws ecr get-login-password` |
| continue | boolean| `bool` |  | |  |  |
| print | boolean| `bool` |  | `true`| If set to false the command will not print the full command with arguments to logs. |  |
| silent | boolean| `bool` |  | | If set to false the command will not print output to logs. |  |

 
