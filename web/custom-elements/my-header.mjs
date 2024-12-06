function MyHeader({ html }) {
  return html` <h1 class="text-2xl font-bold text-[color:--gray-700]">
    <slot></slot>
  </h1>`;
}
