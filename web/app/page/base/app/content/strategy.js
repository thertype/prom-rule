import React, { Component } from 'react'
import { Button, Table, message, Popconfirm, Divider, Input, Icon } from 'antd'
import { getStrategy, addStrategy, updateStrategy, deleteStrategy, getReceiver, addReceiver, updateReceiver, deleteReceiver } from '@apis/strategy'
import Highlighter from 'react-highlight-words'
import CreateEditStrategy from './strategy/create-edit-strategy'
import CreateEditReceiver from './strategy/create-edit-receiver'

export default class Strategy extends Component {
  state = {
    dataSource: [],
    expandData: {},
    filterItem: {
      description: false,
      rule_labels: false,
    },
  }
  currentRow = null
  componentDidMount() {
    this.getList()
    this.expandLoading = false
  }

  getColumnSearchProps = dataIndex => ({
    filterDropdown: ({ setSelectedKeys, selectedKeys, confirm }) => (
      <div style={{ padding: 8 }}>
        <Input
          ref={(node) => {
            this.searchInput = node;
          }}
          placeholder={`Search ${dataIndex}`}
          value={selectedKeys[0]}
          onInput={(e) => { setSelectedKeys(e.target.value ? [e.target.value] : []); this.handleSearch(selectedKeys, confirm, dataIndex) }}
          onBlur={() => this.setState(state => ({
            filterItem: { ...state.filterItem, [dataIndex]: false },
          }))}
          style={{ width: 188, marginBottom: 8, display: 'block' }}
        />
      </div>
    ),
    filterIcon: filtered => (
      <Icon type="search"
        onMouseDown={() => {
          this.setState(state => ({
            filterItem: { ...state.filterItem, [dataIndex]: true },
          })); setTimeout(() => this.searchInput.focus());
        }}
        style={{ color: filtered ? '#1890ff' : undefined }}
      />
    ),
    onFilter: (value, record) => {
      let content
      content = record[dataIndex]
      return content
        .toString()
        .toLowerCase()
        .includes(value.toLowerCase())
    },
    onFilterDropdownVisibleChange: (visible) => {
      if (visible) {
        setTimeout(() => this.searchInput.focus());
      }
    },
    render: text =>
      (this.state.searchedColumn === dataIndex ? (
        <Highlighter
          highlightStyle={{ backgroundColor: '#ffc069', padding: 0 }}
          searchWords={[this.state.searchText]}
          autoEscape
          textToHighlight={text.toString()}
        />
      ) : (
          text
        )),
  })

  handleSearch = (selectedKeys, confirm, dataIndex) => {
    confirm();
    this.setState({
      searchText: selectedKeys[0],
      searchedColumn: dataIndex,
    });
  }

  getList() {
    getStrategy({}, (data) => {
      const obj = {}
      data.forEach((item) => {
        obj[item.id] = []
        if (item.id === 6) {
          obj[6].push({
            date: 6,
          })
        }
      })
      this.setState({
        dataSource: data.sort((a, b) => b.id - a.id).map(item => ({ child: [], ...item })),
        expandData: obj,
      })
    })
  }
  getReceiver = (id) => {
    this.expandLoading = true
    getReceiver({}, { id }, (res) => {
      const { expandData } = this.state
      expandData[id] = res || []
      this.setState({
        expandData,
      })
      this.expandLoading = false
    })
  }

  handleAdd = () => {
    this.createEditStrategy.updateValue()
  }
  handleEdit(record) {
    this.createEditStrategy.updateValue(record)
  }
  handleDelete(record) {
    // eslint-disable-next-line camelcase
    const { id, rule_labels } = record
    deleteStrategy({}, { id }, (res) => {
      // eslint-disable-next-line camelcase
      message.success(`??????${rule_labels}??????`)
      this.getList()
    })
  }
  handleEditReceiver(record) {
    this.createEditReceiver.updateValue({ mode: 'edit', ...record })
  }
  handleDeleteReceiver(record) {
    const { id } = record
    deleteReceiver({}, { id }, (res) => {
      message.success(`??????${id}??????`)
      this.getReceiver(this.currentRow)
    })
  }
  updateStrategy = values => new Promise((resolve) => {
    const { id, ...data } = values
    if (id) {
      updateStrategy(data, { id }, (res) => {
        resolve(true)
        this.getList()
      })
      return
    }
    addStrategy(data, (res) => {
      resolve(true)
      this.getList()
    })
  })
  updateReceiver = value => new Promise((resolve) => {
    const { id, mode, ...data } = value
    if (mode === 'edit') {
      updateReceiver(data, { id }, (res) => {
        resolve(true)
        this.getReceiver(this.currentRow)
      })
      return
    }
    addReceiver(data, { id }, (res) => {
      resolve(true)
      this.getReceiver(id)
    })
  })
  onRefStr(component) {
    this.createEditStrategy = component
  }
  onRefRec(component) {
    this.createEditReceiver = component
  }
  expandedRowRender(recordRow) {
    const { id } = recordRow
    const { expandData } = this.state
    const { expandLoading } = this
    if (!expandLoading) {
      this.getReceiver(id)
    }
    const addStrategyEvent = () => {
      this.createEditReceiver.updateValue({ id, mode: 'create' })
    }
    const columns = [
      {
        title: '???????????????',
        align: 'center',
        dataIndex: 'date',
        render: (text, record) => (
          <span>{record.start_time}~{record.end_time}</span>
        ),
      },
      { title: '????????????', align: 'center', dataIndex: 'start' },
      { title: '????????????', align: 'center', dataIndex: 'period' },
      { title: '????????????', align: 'center', dataIndex: 'user' },
      { title: '?????????', align: 'center', dataIndex: 'duty_group' },
      { title: '???????????????', align: 'center', dataIndex: 'group' },
      { title: 'Filter?????????', align: 'center', dataIndex: 'expression' },
      { title: '????????????', align: 'center', dataIndex: 'method' },
      {
        title: () => (<div>??????<Divider type="vertical" /><a onClick={addStrategyEvent}>??????</a></div>),
        dataIndex: 'operation',
        align: 'center',
        key: 'operation',
        render: (text, record) => (
          <span>
            <a onClick={() => { this.currentRow = id; this.handleEditReceiver(record) }}>??????</a>
            {/* <Divider type="vertical" /> */}
            <Popconfirm
              title="???????????????????"
              onConfirm={() => { this.currentRow = id; this.handleDeleteReceiver(record) }}
              okText="Yes"
              cancelText="No"
            >
              <a href="#">??????</a>
            </Popconfirm>
          </span>
        ),
      },
    ];
    return <Table columns={columns} dataSource={expandData[id]} pagination={false} rowKey="id" />
  }
  render() {
    const { dataSource } = this.state
    const columns = [
      {
        title: '??????', align: 'center', dataIndex: 'id', key: 'id', sorter: (a, b) => a.id - b.id,
      },
      {
        title: '??????',
        align: 'center',
        dataIndex: 'description',
        key: 'description',
        ...this.getColumnSearchProps('description'),
        filterDropdownVisible: this.state.filterItem.description,
      },
      {
        title: '??????',
        align: 'center',
        dataIndex: 'rule_labels',
        key: 'rule_labels',
        ...this.getColumnSearchProps('rule_labels'),
        filterDropdownVisible: this.state.filterItem.rule_labels,
      },
      {
        title: '??????',
        align: 'center',
        key: 'action',
        render: (text, record, index) => (
          <span>
            <a onClick={() => this.handleEdit(record)}>??????</a>
            <Divider type="vertical" />
            <Popconfirm
              title="???????????????????"
              onConfirm={() => this.handleDelete(record)}
              okText="Yes"
              cancelText="No"
            >
              <a href="#">??????</a>
            </Popconfirm>
          </span>
        ),
      },
    ]
    return (
      <div>
        <div id="top-section">
          <Button type="primary" onClick={this.handleAdd}>??????</Button>
        </div>
        <Table dataSource={dataSource} expandedRowRender={record => this.expandedRowRender(record)} columns={columns} rowKey="id" />
        <CreateEditStrategy OnRef={c => this.onRefStr(c)} onSubmit={this.updateStrategy} />
        <CreateEditReceiver OnRef={c => this.onRefRec(c)} onSubmit={this.updateReceiver} />
      </div>
    )
  }
}
