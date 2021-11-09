import React from 'react';
import { render, screen } from '@testing-library/react';
import App from './components/App';
import '@testing-library/jest-dom';

test('renders biography paragraph', () => {
  render(<App />);
  const linkElement = screen.getByText(/Hi! My name is Andrew Willette. I am a software developer based in Kansas City, Kansas, where I work primarily with cloud-hosted security protocol services, both their backend implementations and frontend management UIs. I like playing violin and host some recordings on my site here./i);
  expect(linkElement).toBeInTheDocument();
});
