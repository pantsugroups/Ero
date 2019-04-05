# -*- coding: utf-8 -*-
import ero_api
import config

app = ero_api.create_app(config)
app.run(host=config.HOSTNAME,
        port=config.PORT,
        debug=config.DEBUG)
