import 'babel-polyfill'
import React from 'react'
import ReactDOM from 'react-dom'
import { Provider } from 'react-redux'
import '@/configs/config'
import Routes from '@/configs/router.config.js'
import configure from '@/middleware/configureStore'

const store = configure({ })
ReactDOM.render(
  <Provider store={store}>
    <Routes />
  </Provider>,
  document.getElementById('root'),
)
