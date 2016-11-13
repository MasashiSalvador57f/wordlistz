## database settings in development environment
```
CREATE DATABASE wordzlist_development;
GRANT ALL ON wordzlist_development.* TO wordzlist_user IDENTIFIED BY "wordzlist_dev";
```

```
% goose status                                                                                                                                                                               (git)-[master]
goose: status for environment 'development'
    Applied At                  Migration
    =======================================
```