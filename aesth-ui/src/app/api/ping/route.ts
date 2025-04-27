import { cookies } from "next/headers";
import { NextRequest, NextResponse } from "next/server";

const NEXT_PUBLIC_SERVER_URL = process.env.NEXT_PUBLIC_SERVER_URL;
export async function GET(req: NextRequest) {
  try {
    const cookieStore = await cookies();
    const token = cookieStore.get("access_token")?.value;
  
    if (!token) {
      console.warn("No access token found in cookies");
      return NextResponse.redirect(new URL("/auth", req.url));
    }
  
    const serverResponse = await fetch(`${NEXT_PUBLIC_SERVER_URL}/api/ping`, {
      method: "GET",
      headers: {
        "Authorization": `Bearer ${token}`,
      },
    });
  
    if (serverResponse.status === 401) {
      console.warn("Unauthorized access, redirecting to /auth");
      return NextResponse.redirect(new URL("/auth", req.url));
    }
  
    return NextResponse.json(await serverResponse.json());
  } catch (e: any) {
    console.error("Unexpected error:", e);
    return NextResponse.json(
      { error: "Internal server error", details: e.message },
      { status: 500 }
    );
  }
}
