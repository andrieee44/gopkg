#include <stdlib.h>
#include <stddef.h>
#include <linux/uinput.h>

unsigned int wrapUI_GET_SYSNAME(int len) {
	return UI_GET_SYSNAME(len);
}

size_t* layout_uinput_ff_upload(size_t* count) {
	*count = 5;

	size_t* layout = malloc(sizeof(size_t) * (*count));
	if (!layout) {
		*count = 0;

		return NULL;
	}

	layout[0] = offsetof(struct uinput_ff_upload, request_id);
	layout[1] = offsetof(struct uinput_ff_upload, retval);
	layout[2] = offsetof(struct uinput_ff_upload, effect);
	layout[3] = offsetof(struct uinput_ff_upload, old);
	layout[4] = sizeof(struct uinput_ff_upload);

	return layout;
}

size_t* layout_uinput_ff_erase(size_t* count) {
	*count = 4;

	size_t* layout = malloc(sizeof(size_t) * (*count));
	if (!layout) {
		*count = 0;

		return NULL;
	}

	layout[0] = offsetof(struct uinput_ff_erase, request_id);
	layout[1] = offsetof(struct uinput_ff_erase, retval);
	layout[2] = offsetof(struct uinput_ff_erase, effect_id);
	layout[3] = sizeof(struct uinput_ff_erase);

	return layout;
}

size_t* layout_uinput_setup(size_t* count) {
	*count = 4;

	size_t* layout = malloc(sizeof(size_t) * (*count));
	if (!layout) {
		*count = 0;

		return NULL;
	}

	layout[0] = offsetof(struct uinput_setup, id);
	layout[1] = offsetof(struct uinput_setup, name);
	layout[2] = offsetof(struct uinput_setup, ff_effects_max);
	layout[3] = sizeof(struct uinput_setup);

	return layout;
}

size_t* layout_uinput_abs_setup(size_t* count) {
	*count = 3;

	size_t* layout = malloc(sizeof(size_t) * (*count));
	if (!layout) {
		*count = 0;

		return NULL;
	}

	layout[0] = offsetof(struct uinput_abs_setup, code);
	layout[1] = offsetof(struct uinput_abs_setup, absinfo);
	layout[2] = sizeof(struct uinput_abs_setup);

	return layout;
}

size_t* layout_uinput_user_dev(size_t* count) {
	*count = 8;

	size_t* layout = malloc(sizeof(size_t) * (*count));
	if (!layout) {
		*count = 0;

		return NULL;
	}

	layout[0] = offsetof(struct uinput_user_dev, name);
	layout[1] = offsetof(struct uinput_user_dev, id);
	layout[2] = offsetof(struct uinput_user_dev, ff_effects_max);
	layout[3] = offsetof(struct uinput_user_dev, absmax);
	layout[4] = offsetof(struct uinput_user_dev, absmin);
	layout[5] = offsetof(struct uinput_user_dev, absfuzz);
	layout[6] = offsetof(struct uinput_user_dev, absflat);
	layout[7] = sizeof(struct uinput_user_dev);

	return layout;
}
