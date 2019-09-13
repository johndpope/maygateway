import styled from 'styled-components'
import { NavLink } from 'react-router-dom'

const Wrapper = styled.div`
  width: 240px;
  box-shadow: inset -2px 2px 15px rgba(0,0,0,.08);
`

const LogoWrapper = styled.div`
  padding: 10px 20px;

  img {
    width: 100%;
  }
`

const List = styled.div`
  padding: 0 15px;

  .selected {
    background: #304ec2;
    color: #fff;
    border-radius: 10px;

    svg {
      color: #fff;
    }
  }

`

const ListItem = styled(NavLink)`
  padding: 10px;
  margin: 5px 0;
  color: #79797d;
  display: block;
  text-decoration: none;

  svg {
    margin-right: 10px;
    color: #98979c;
  }
`

export default {
  Wrapper,
  LogoWrapper,
  List,
  ListItem,
}