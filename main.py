# -*- coding: utf-8 -*-
import ero_api
import config

app = ero_api.create_app(config)

if __name__ == "__main__":
    app.run(host=config.HOSTNAME,
            port=config.PORT,
            debug=True)
