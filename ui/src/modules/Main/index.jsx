import React from 'react'
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'
import { GlobalStyle } from 'core/styled/global'
import Apis from 'modules/Apis'
import Home from 'modules/Home'
import Sidebar from './Sidebar'
import Navbar from './Navbar'
import Styled from './styled'

const Main = () => {
  return (
    <Router>
      <Styled.Wrapper>
        <Styled.Content>
          <Sidebar />
          <Styled.Main>
            <Switch>
              <Route exact path="/" component={Home} />
              <Route path="/apis" component={Apis} />
            </Switch>
          </Styled.Main>
        </Styled.Content>
        <GlobalStyle />
      </Styled.Wrapper>
    </Router>
  )
}

export default Main