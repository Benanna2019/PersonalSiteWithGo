function IndexPage({ html, state }) {
  const { store } = state;
  const { limit, offset, total } = store;

  return html`
    <site-layout>
      <main>
        <blog-posts></blog-posts>
        <blog-pagination
          limit="${limit}"
          offset="${offset}"
          total="${total}"
          class="pb3 pb5-lg"
        ></blog-pagination>
      </main>
    </site-layout>
  `;
}
// <my-h-card class="hidden"></my-h-card>
