JSONSelector
============

Jsonelector allows you to find matching nodes in a JSON document using
a query language modeled off of CSS selectors -- 
`JSONSelect <http://jsonselect.org/>`_.

Installation
------------

Install using ``go get``::

    go get github.com/coddingtonbear/jsonselector

Quick Start
-----------

Command::

    curl https://raw.githubusercontent.com/lloyd/JSONSelectTests/master/level_1/basic.json | jsonselector .favoriteColor

Output::

    "yellow"

Usage
-----

Jsonselector accepts a piped-in JSON document and a single argument holding the selectors
you'd like to retrieve; the following two examples are equivalent and will return all
nodes in the JSON document ``some_json_document`` matching the selector ``.my_selector``::

    cat some_json_document | jsonselector .my_selector
    jsonselector < some_json_document .my_selector
