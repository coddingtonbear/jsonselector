.. image:: https://travis-ci.org/coddingtonbear/go-jsonselect.png?branch=master
   :target: https://travis-ci.org/coddingtonbear/go-jsonselect

Jsonselector allows you to find matching nodes in a JSON document using
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


Acknowledgements
----------------

The chains of inspiration are long;
this project was inspired by `@EricChang <https://github.com/EricChiang>`_'s `pup <https://github.com/EricChiang/pup>`_
-- a commandline HTML parser,
and is really not much more than a thin wrapper around
the `go-jsonselect <https://github.com/coddingtonbear/go-jsonselect>`_ library
I cobbled together a few months ago,
which was itself inspired by
a `Python implementation of the JSONSelect query language <https://github.com/mwhooker/jsonselect>`_
by `@mwhooker <https://github.com/mwhooker>`_.
