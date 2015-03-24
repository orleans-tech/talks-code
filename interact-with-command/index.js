#!/usr/bin/env node

var CommandAsker    = require('command-asker'),
    when            = require('when'),
    validators      = require('./lib/validators');

var a = new CommandAsker([
    { key: 'firstName', ask: 'Quel est ton prénom' },
    { key: 'lastName',  ask: 'Quel est ton nom',    required: true },
    { key: 'age',       ask: 'Quel âge as-tu',      validators: [validators.isAdult] }
]);

// Launch the prompt command
a.ask(function(res) {
    console.log('Mon nom est ' + res.firstName + ' ' + res.lastName + ' et j\'ai ' + res.age + ' ans !');

    // Close the prompt command
    a.close();
});