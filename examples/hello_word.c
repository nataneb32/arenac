#include <arenac/ctx.h>

void main() {
	arenac_ctx_t ctx;
	arenac_ctx_init(&ctx);

	arenac_printf(&ctx, "Hello, world!");

	arenac_ctx_destroy(&ctx);
}
