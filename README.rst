zang-go
==========

This golang package is an open source tool built to simplify interaction with
the `Zang <http://www.zang.io>`_ telephony platform. Zang Cloud makes adding voice
and SMS to applications fun and easy.

For more information about Zang, please visit:
`Zang Cloud <https://www.zang.io/products/cloud>`_

To read the official documentation visit `Zang Docs <http://docs.zang.io>`_.


Installation
============

Clone the repo, and install via ``go get``:

.. code-block:: bash
    $ go get -u github.com/zang-cloud/zang-go
.. code-block:: bash

Usage
======

REST
----

See the `Zang REST API documentation <http://docs.zang.io/docs/overview>`_
for more information.

Send SMS Example
----------------

.. code-block:: golang


.. code-block:: golang

InboundXML
==========

InboundXML is an XML dialect which enables you to control phone call flow.
For more information please visit the `Zang InboundXML documentation
<http://docs.zang.io/docs/inboundxml-overview>`_.

<Say> Example
-------------

.. code-block:: golang

  ixml, err := New(Response{Say: &Say{
    Voice: "female",
    Value: "Welcome to Zang!",
    Loop: 3,
  }})

  fmt.Print(ixml)

.. code-block:: golang

will render

.. code-block:: xml

    <?xml version="1.0" encoding="UTF-8" standalone="yes"?>
    <Response>
        <Say loop="3" voice="female" language="en">Welcome to Zang!</Say>
    </Response>

.. code-block:: xml