import React from 'react'
import Styled from './styled'
import { Icon } from 'core/components'
import Logo from 'core/assets/images/maylogo.png'

const Sidebar = () => {

  const renderListItem = (path, icon, text, exact) => (
    <Styled.ListItem exact to={path} activeClassName="selected"><Icon icon={icon} />{text}</Styled.ListItem>
  )

  return (
    <Styled.Wrapper>
      <Styled.LogoWrapper>
        <img src={Logo} />
      </Styled.LogoWrapper>
      <Styled.List>
        { renderListItem('/', 'home', 'Home', true) }
        { renderListItem('/apis', 'cloud', 'Apis') }
      </Styled.List>
    </Styled.Wrapper>
  )
}

export default Sidebar