import React, { Component } from 'react';
import { Menu, Icon } from 'antd';
import { Switch, withRouter } from 'react-router-dom'
import { getMethod } from '@apis/login';

// 路由对应 sider 的匹配规则
const regRouter = new Map([
  [/^\/rules$/, 'rules'],
  [/^\/promethus$/, 'promethus'],
  [/^\/strategy$/, 'strategy'],
  [/^\/rulegroup$/, 'rulegroup'],
  [/^\/alerts$/, 'alerts'],
  [/^\/alerts_confirm\/?[\d]+$/, 'alerts_confirm'],
  [/^\/alerts_confirm$/, 'alerts_confirm'],
  [/^\/group$/, 'group'],
  [/^\/maintain$/, 'maintain'],
  [/^\/user$/, 'user'],
])

@withRouter
export default class Sider extends Component {
  constructor(props) {
    super(props);
    this.menuClick = this.menuClick.bind(this);
  }
  state = {
    usermanage: true,
    collapsed: false,
    selectedKeys: [],
  }
  activeKey = undefined
  // event
  toggleCollapsed = () => {
    this.setState({
      collapsed: !this.state.collapsed,
    });
  }
  menuClick = (e) => {
    const { history } = this.props
    history.push(`/${e.key}`);
  }
  setMenuActive() {
    const { pathname } = this.props.history.location
    let activeKey
    if (pathname) {
      [...regRouter].some(([reg, path]) => {
        if (reg.test(pathname)) {
          activeKey = path
          return true
        }
        return false
      })
      if (activeKey !== this.activeKey) {
        this.activeKey = activeKey
        this.setState({
          selectedKeys: [activeKey || undefined],
        })
      }
    }
  }
  componentWillMount() {
    getMethod({}, (res) => {
      if (res === 'local') {
        this.setState({
          usermanage: false,
        });
      }
    });
  }
  componentDidMount() {
    this.setMenuActive()
  }
  componentDidUpdate() {
    this.setMenuActive()
  }
  render() {
    const { selectedKeys } = this.state
    return (
      <Switch>
        <div id="sidenav" style={{ width: 220 }}>
          <Menu
            defaultSelectedKeys={['1']}
            defaultOpenKeys={['sub1']}
            mode="inline"
            theme="dark"
            selectedKeys={selectedKeys}
            inlineCollapsed={this.state.collapsed}
            onClick={this.menuClick}
          >
            <Menu.Item key="rules">
              <Icon type="audit" />
              <span>告警规则管理</span>
            </Menu.Item>
            <Menu.Item key="promethus">
              <Icon type="desktop" />
              <span>数据源管理</span>
            </Menu.Item>
            <Menu.Item key="rulegroup">
              <Icon type="contacts" />
              <span>告警模板管理</span>
            </Menu.Item>
            <Menu.Item key="strategy">
              <Icon type="gateway" />
              <span>告警计划管理</span>
            </Menu.Item>
          </Menu>
        </div>
      </Switch>
    );
  }
}
