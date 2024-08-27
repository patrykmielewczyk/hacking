use MIME::Base64;

my $encoded = "aWQ=";  # base64 coded 'id'
my $decoded = decode_base64($encoded);

system($decoded);
