Feature: gpg function test

Background:
* url demoBaseUrl
* def secrets = read('secrets/secrets.json')

Scenario: simple test

    Given path '/'
    Given header Direktiv-ActionID = 'development'
    Given header Direktiv-Tempdir = '/tmp'
    And request 
    """
    { 
      "access-key": '#(secrets.key)',
      "secret-key": '#(secrets.secret)',
      "region": '#(secrets.region)',
      "commands": [
        {
          "command": "aws ecr get-login-password",
          "silent": true,
          "print": false,
        }
      ]
    }
    """
    When method post
    Then status 200
    And match $ == 
    """
    {
    "aws": [
    {
      "result": "#notnull",
      "success": true
    }
    ]
    }
    """

Scenario: wrong account

    Given path '/'
    Given header Direktiv-ActionID = 'development'
    Given header Direktiv-Tempdir = '/tmp'
    And request 
    """
    { 
      "access-key": 'NOACCESSKEY',
      "secret-key": 'NOSECRET',
      "region": 'NOREGION',
      "commands": [
        {
          "command": "aws ecr get-login-password",
          "silent": true,
          "print": false,
        }
      ]
    }
    """
    When method post
    Then status 500
    And def temp = response['Direktiv-Errorcode']
    And match header Direktiv-Errorcode == 'io.direktiv.command.error'

Scenario: continue on error

    Given path '/'
    Given header Direktiv-ActionID = 'development'
    Given header Direktiv-Tempdir = '/tmp'
    And request 
    """
    { 
      "access-key": '#(secrets.key)',
      "secret-key": '#(secrets.secret)',
      "region": '#(secrets.region)',
      "commands": [
        {
        "command": "aws does not exist",
        "continue": true
        },
        {
          "command": "aws ecr get-login-password",
          "silent": true,
          "print": false,
        }
      ],
    }
    """
    When method post
    Then status 200

Scenario: stop on error

    Given path '/'
    Given header Direktiv-ActionID = 'development'
    Given header Direktiv-Tempdir = '/tmp'
    And request 
    """
    { 
      "access-key": '#(secrets.key)',
      "secret-key": '#(secrets.secret)',
      "region": '#(secrets.region)',
      "commands": [
        {
        "command": "aws does not exist",
        "continue": false
        },
        {
          "command": "aws ecr get-login-password",
          "silent": true,
          "print": false,
        }
      ]
    }
    """
    When method post
    Then status 500
