# objects
- country
    - id uuid
    - bandName string
    - songName string
    - flag  string
    - votes []vote
    - comments  []comment
    - participating boolean

- user
    - id uuid
    - authLvl enum
    - name string
    - comments []comment
    - icon string
    - votes []vote

- comment
    - id uuid
    - userId uuid
    - countryId uuid
    - text text

- authLvl
    - admin
    - user

- vote
    - id uuid
    - countryId uuid
    - userId uuid
    - value int
