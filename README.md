# genericbot

## Tasks
### Database
- [ ] Create `aliases`
- [ ] Migrate `auditlog`
- [ ] Migrate `bans`
  - what the fuck timestamp format does the C# Mongo driver use???
- [x] Migrate `config`
- [ ] Migrate `customCommands` to `commands`
- [ ] Migrate `giveaways`
- [x] Migrate `quotes` (no work needed)
- [ ] Migrate `users`
  - _id field needs to be converted
  - don't bother migrating LastPointsAdded field (timestamp)
