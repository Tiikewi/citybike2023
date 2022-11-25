import { CustomLink } from '../components/NavbarLink';
import '../styles/navbar.css';



export const Navbar = (): JSX.Element | null => {
  return (
    <nav className="nav">
      <ul className="container">
          <CustomLink to="/">Home</CustomLink>
          <CustomLink to="/test">Test</CustomLink>
      </ul>
    </nav>
  );
};