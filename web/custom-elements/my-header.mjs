function MyHeader({ html }) {
  return html` <h1 class="text-2xl font-bold text-red-500">
    <slot></slot>
  </h1>`;
}
