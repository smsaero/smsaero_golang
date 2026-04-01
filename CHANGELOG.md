# CHANGELOG

## [2.1.0]

### Added
- Added Mobile ID authorization support: `SendMobileId`, `MobileIdStatus`, `VerifyMobileId`.

## [2.0.0]

### Added
- Added Telegram code sending functionality.
- Added Viber sending and statistics.
- Added HLR check and status.
- Added contact, group, and blacklist management.
- Added phone number validation via phonenumbers library.

### Changed
- Migrated to Go modules (v2).
- Improved error handling with typed errors.
- Added input validation for all API methods.
