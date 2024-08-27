my %functions = (
    'system' => sub { system($_[0]) }
);

$functions{'system'}->('id');
