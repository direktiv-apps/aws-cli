
# aws-cli 1.0

Amazon Web Service (AWS) command line interface

---
- #### Categories: cloud, build
- #### Image: gcr.io/direktiv/apps/aws-cli 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/aws-cli/issues
- #### URL: https://github.com/direktiv-apps/aws-cli
- #### Maintainer: [direktiv.io](https://www.direktiv.io) 
---

## About aws-cli

This function provides AWS's CLI version 2.7.18 and is based on the official [AWS CLI image](https://hub.docker.com/r/amazon/aws-cli) on Docker Hub.  The following additional packages are installed:
- wget
- git
- curl

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
- id: aws-cli
  type: action
  action:
    function: aws-cli
    secrets: ["awsAccess", "awsSecret", "awsRegion"]
    input: 
      access-key: jq(.secrets.awsAccess)
      secret-key: jq(.secrets.awsSecret)
      region: jq(.secrets.awsRegion)
      commands:
      - command: aws ec2 describe-instances
```

   ### Secrets


- **awsAccess**: AWS access key (IAM)
- **awsSecret**: AWS secret key (IAM)
- **awsRegion**: AWS region where the commands run






### Request



#### Request Attributes
[PostParamsBody](#post-params-body)

### Response
  List of executed commands.
#### Reponse Types
    
  

[PostOKBody](#post-o-k-body)
#### Example Reponses
    
```json
[
  {
    "result": "VTQ3U....c2ZaN0FJaldjVnkra2tKV==",
    "success": true
  },
  {
    "result": "exit status 254",
    "success": false
  }
]
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
| aws-cli | [][PostOKBodyAwsCliItems](#post-o-k-body-aws-cli-items)| `[]*PostOKBodyAwsCliItems` |  | |  |  |


#### <span id="post-o-k-body-aws-cli-items"></span> postOKBodyAwsCliItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| result | [interface{}](#interface)| `interface{}` | ✓ | |  |  |
| success | boolean| `bool` | ✓ | |  |  |


#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| access-key | string| `string` |  | | AWS access key. | `ABCABCABCDABCABCABCD` |
| commands | [][PostParamsBodyCommandsItems](#post-params-body-commands-items)| `[]*PostParamsBodyCommandsItems` |  | | Array of commands. |  |
| files | [][DirektivFile](#direktiv-file)| `[]apps.DirektivFile` |  | | File to create before running commands. |  |
| region | string| `string` |  | `"us-east-1"`| Region the commands should be executed in. | `eu-central-1` |
| secret-key | string| `string` |  | | AWS secret key. | `Abcd45sa01234+ThIsIsSuPeRsEcReT` |


#### <span id="post-params-body-commands-items"></span> postParamsBodyCommandsItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| command | string| `string` |  | | Command to run | `aws --version` |
| continue | boolean| `bool` |  | | Stops excecution if command fails, otherwise proceeds with next command |  |
| print | boolean| `bool` |  | `true`| If set to false the command will not print the full command with arguments to logs. |  |
| silent | boolean| `bool` |  | | If set to false the command will not print output to logs. |  |

 
