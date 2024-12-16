function BlogContainer({ html }) {
  return html`
    <style>
      :host {
        display: block;
        max-width: 90vw;
        margin-inline: auto;
      }

      @media screen and (min-width: 48em) {
        :host {
          max-width: 82ch;
        }
      }
    </style>

    <div class="dark:text-white text-black">
      <slot></slot>
    </div>
  `;
}
