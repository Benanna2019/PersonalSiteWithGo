function SiteHeader({ html }) {
  return html`
    <style>
      h1,
      p {
        text-align: var(--align-heading);
      }
    </style>
    <header>
      <h1 class="text3 font-heading font-bold tracking-1 pbs4 pbe0">
        <a href="/" class="no-underline"> Ben Patton </a>
      </h1>
      <p class="font-body text0 pbe4">
        Thoughts. Technical tips. And other things.
      </p>
    </header>
  `;
}
