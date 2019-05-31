import React from 'react';
import {Breadcrumb, Icon, Layout, Menu} from 'antd';
import './App.less';
import AudioPlayer from "./common/AudioPlayer";

const {Header, Content, Footer, Sider} = Layout;
const {SubMenu} = Menu;

class App extends React.Component<any, any> {
  state = {
    collapsed: false,
  };

  onCollapse = (collapsed: any) => {
    console.log(collapsed);
    this.setState({collapsed});
  };

  render() {
    return (
        <Layout style={{minHeight: '100vh'}}>
          <Sider collapsible collapsed={this.state.collapsed} onCollapse={this.onCollapse}>
            <div className="logo"/>
            <Menu theme="dark" defaultSelectedKeys={['1']} mode="inline">
              <Menu.Item key="1">
                <Icon type="pie-chart"/>
                <span>Home</span>
              </Menu.Item>
              <Menu.Item key="2">
                <Icon type="desktop"/>
                <span>Music</span>
              </Menu.Item>
              <SubMenu
                  key="sub1"
                  title={
                    <span>
                                  <Icon type="user"/>
                                  <span>Artist</span>
                                </span>
                  }
              >
                <Menu.Item key="3">Tommy</Menu.Item>
                <Menu.Item key="4">Joshua</Menu.Item>
                <Menu.Item key="5">Cheever</Menu.Item>
              </SubMenu>
              <SubMenu
                  key="sub2"
                  title={
                    <span>
                                  <Icon type="team"/>
                                  <span>Followers</span>
                                </span>
                  }
              >
                <Menu.Item key="6">Team 1</Menu.Item>
                <Menu.Item key="8">Team 2</Menu.Item>
              </SubMenu>
              <Menu.Item key="9">
                <Icon type="file"/>
                <span>Upload Music</span>
              </Menu.Item>
            </Menu>
          </Sider>
          <Layout>
            <Header style={{background: '#fff', padding: 0}}/>
            <Content style={{margin: '0 16px'}}>
              <Breadcrumb style={{margin: '16px 0'}}>
                <Breadcrumb.Item>Music</Breadcrumb.Item>
                <Breadcrumb.Item>Good Things Fall Apart</Breadcrumb.Item>
              </Breadcrumb>
              <div style={{padding: 24, background: '#fff', minHeight: 360}}>
                Artist: Jon Bellion & Illenium
              </div>
            </Content>
            <Footer style={{textAlign: 'center', padding: 0, height: 100}}>
              <AudioPlayer/>
            </Footer>
          </Layout>
        </Layout>
    );
  }
}

export default App;
