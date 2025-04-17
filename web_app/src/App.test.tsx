import { render, screen } from '@testing-library/react';
import App from './App';

test('renders "container" id', () => {
  render(<App />);
  const container = screen.getByTestId(/container/i);
  expect(container).toBeInTheDocument();
});
