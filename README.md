# notification_service

This is a project to address the implementation of a rate limiter for an email notification system.

The system must reject requests that exceed the limit. Some sample notification types and rate limit rules include:

- Status: No more than 2 per minute for each recipient.
- News: No more than 1 per day for each recipient.
- Marketing: No more than 3 per hour for each recipient.

*These are just examples; the system may have multiple rate limit rules.*