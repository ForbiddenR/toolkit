package transport

import "context"

type msgSNKey struct{}

func MsgSNFromCtx(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if v, ok := ctx.Value(msgSNKey{}).(string); ok {
		return v
	}
	return ""
}

func CtxWithMsgSN(ctx context.Context, msgSN string) context.Context {
	return context.WithValue(ctx, msgSNKey{}, msgSN)
}

type msgIndentifyKey struct{}

func MsgIDFromCtx(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if v, ok := ctx.Value(msgIndentifyKey{}).(string); ok {
		return v
	}
	return ""
}

func CtxWithMsgID(ctx context.Context, msgID string) context.Context {
	return context.WithValue(ctx, msgIndentifyKey{}, msgID)
}

type traceIdKey struct{}

func TraceIDFromCtx(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if v, ok := ctx.Value(traceIdKey{}).(string); ok {
		return v
	}
	return ""
}

func CtxWithTraceID(ctx context.Context, traceId string) context.Context {
	return context.WithValue(ctx, traceIdKey{}, traceId)
}

type equipmentIdKey struct{}

func EquipmentIdFromCtx(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if v, ok := ctx.Value(equipmentIdKey{}).(string); ok {
		return v
	}
	return ""
}

func CtxWithEquipmentId(ctx context.Context, equipmentId string) context.Context {
	return context.WithValue(ctx, equipmentIdKey{}, equipmentId)
}
