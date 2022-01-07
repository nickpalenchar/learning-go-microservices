

### Dependency Injection

A strategy where objects implement others through reference rather than directly coded, so that they can be
easilly swapped out in the case of

### Graceful Timeouts

An important concepts in servers. If a server just abruptly stops (from an error or when doing an upgrade
for example), connected users will not be able to finish their work.

A graceful timeout will be able to allow for all in-progress clients to complete their requests before
shutting down.