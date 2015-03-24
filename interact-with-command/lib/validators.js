'use strict';

var when = require('when');

var isAdult = function(value) {
    return when.promise(function(resolve, reject) {
        var num = Number(value);
        if (isNaN(num) || num < 18) {
            return reject({
                'name'      : 'not_adult',
                'message'   : 'vous devez être majeur'
            });
        }
        return resolve();
    });
};
module.exports = {
    isAdult: isAdult
};