
Feature: Basic

# The secrects can be used in the payload with the following syntax #(mysecretname)
Background:
* def awsAccess = karate.properties['awsAccess']
* def awsSecret = karate.properties['awsSecret']
* def awsRegion = karate.properties['awsRegion']


Scenario: version

	Given url karate.properties['testURL']

	And path '/'
	And header Direktiv-ActionID = 'development'
	And header Direktiv-TempDir = '/tmp'
	And request
	"""
	{	
		"access-key": #(awsAccess),
		"secret-key": #(awsSecret),
		"awsRegion": #(awsRegion),
		"commands": [
		{
			"command": "aws --version",
			"silent": true,
			"print": false,
		}
		]
	}
	"""
	When method POST
	Then status 200
		And match $ ==
	"""
	{
	"aws-cli": [
	{
		"result": "#notnull",
		"success": true
	}
	]
	}
	"""

Scenario: listvm

	Given url karate.properties['testURL']

	And path '/'
	And header Direktiv-ActionID = 'development'
	And header Direktiv-TempDir = '/tmp'
	And request
	"""
	{	
		"access-key": #(awsAccess),
		"secret-key": #(awsSecret),
		"awsRegion": #(awsRegion),
		"commands": [
		{
			"command": "aws ec2 describe-instances",
			"silent": true,
			"print": false,
		}
		]
	}
	"""
	When method POST
	Then status 200
		And match $ ==
	"""
	{
	"aws-cli": [
	{
		"result": "#notnull",
		"success": true
	}
	]
	}
	"""
