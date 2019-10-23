# gmail-api-sendmail
Emulate sendmail via Gmail API

This can be used as a sendmail-compatible binary that will send emails via the GMail API, possibly with queue.

Use the gmail-sendmail binary as "sendmail".
If a ~/.gmail-queue folder exists, it will queue up any sent emails there, if not it will immediately send them.

If you have queued any emails, you can run gmail-sendqueued to process the queue and send the emails out.
