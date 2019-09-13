import styled from 'styled-components'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'

const Wrapper = styled(FontAwesomeIcon)`
  color: ${({ color }) => color || '#000'}
`

export default {
  Wrapper,
}