import { NextRequest, NextResponse } from "next/server";

const NEXT_PUBLIC_SERVER_URL = process.env.NEXT_PUBLIC_SERVER_URL;
export async function POST(req: NextRequest) {
  try {
    const { email, password } = await req.json();

    const serverResponse = await fetch(`${NEXT_PUBLIC_SERVER_URL}/auth`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email, password }),
    });
  
    if (serverResponse.status === 401) {
      console.warn("Unauthorized access, redirecting to /auth");
      return NextResponse.redirect(new URL("/auth", req.url));
    }
  
    const { token } = await serverResponse.json();
  
    const response = NextResponse.json({ message: "Logged in" });
    response.cookies.set("access_token", token, {
      httpOnly: true,
      secure: process.env.NODE_ENV === "production",
      path: "/",
      sameSite: "strict",
      maxAge: 60 * 60 * 24,
    });
  
    return response;
  } catch (e: any) {
    console.error("Unexpected error:", e);
    return NextResponse.json(
      { error: "Internal server error", details: e.message },
      { status: 500 }
    );
  }
}